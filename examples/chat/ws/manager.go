package ws

import (
	"fmt"
	"github.com/puzpuzpuz/xsync/v3"
	"net/http"
)

type EventType string

const (
	ConnectedEvent    EventType = "connected"
	DisconnectedEvent EventType = "disconnected"
	MessageEvent      EventType = "message"
)

type SocketEvent struct {
	Id      string
	RoomId  string
	Type    EventType
	Payload map[string]any
}

type CloseEvent struct {
	Code   int
	Reason string
}

type SocketConnection struct {
	Id     string
	Writer http.ResponseWriter
	RoomId string
	Done   chan CloseEvent
	Flush  chan bool
}

type SocketManager struct {
	sockets   *xsync.MapOf[string, *xsync.MapOf[string, SocketConnection]]
	idToRoom  *xsync.MapOf[string, string]
	listeners []chan SocketEvent
}

func NewSocketManager() *SocketManager {
	return &SocketManager{
		sockets:  xsync.NewMapOf[string, *xsync.MapOf[string, SocketConnection]](),
		idToRoom: xsync.NewMapOf[string, string](),
	}
}

func (manager *SocketManager) ForEachSocket(roomId string, cb func(conn SocketConnection)) {
	sockets, ok := manager.sockets.Load(roomId)
	if !ok {
		return
	}
	sockets.Range(func(id string, conn SocketConnection) bool {
		cb(conn)
		return true
	})
}

func (manager *SocketManager) Listen(listener chan SocketEvent) {
	if manager.listeners == nil {
		manager.listeners = make([]chan SocketEvent, 0)
	}
	manager.listeners = append(manager.listeners, listener)
}

func (manager *SocketManager) dispatch(event SocketEvent) {
	for _, listener := range manager.listeners {
		listener <- event
	}
}

func (manager *SocketManager) OnMessage(id string, message map[string]any) {
	socket := manager.Get(id)
	if socket == nil {
		return
	}
	manager.dispatch(SocketEvent{
		Id:      id,
		Type:    MessageEvent,
		Payload: message,
		RoomId:  socket.RoomId,
	})
}

func (manager *SocketManager) Add(roomId string, id string, writer http.ResponseWriter, done chan CloseEvent, flush chan bool) {
	manager.idToRoom.Store(id, roomId)

	sockets, ok := manager.sockets.LoadOrCompute(roomId, func() *xsync.MapOf[string, SocketConnection] {
		return xsync.NewMapOf[string, SocketConnection]()
	})

	sockets.Store(id, SocketConnection{
		Id:     id,
		Writer: writer,
		RoomId: roomId,
		Done:   done,
		Flush:  flush,
	})

	s, ok := sockets.Load(id)
	if !ok {
		return
	}

	manager.dispatch(SocketEvent{
		Id:      s.Id,
		Type:    ConnectedEvent,
		RoomId:  s.RoomId,
		Payload: map[string]any{},
	})
}

func (manager *SocketManager) OnClose(id string) {
	socket := manager.Get(id)
	if socket == nil {
		return
	}
	manager.dispatch(SocketEvent{
		Id:      id,
		Type:    DisconnectedEvent,
		RoomId:  socket.RoomId,
		Payload: map[string]any{},
	})
	manager.sockets.Delete(id)
}

func (manager *SocketManager) CloseWithError(id string, code int, message string) {
	conn := manager.Get(id)
	if conn != nil {
		go manager.OnClose(id)
		conn.Done <- CloseEvent{
			Code:   code,
			Reason: message,
		}
	}
}

func (manager *SocketManager) Disconnect(id string) {
	conn := manager.Get(id)
	if conn != nil {
		go manager.OnClose(id)
		conn.Done <- CloseEvent{
			Code:   -1,
			Reason: "",
		}
	}
}

func (manager *SocketManager) Get(id string) *SocketConnection {
	roomId, ok := manager.idToRoom.Load(id)
	if !ok {
		return nil
	}
	sockets, ok := manager.sockets.Load(roomId)
	if !ok {
		return nil
	}
	conn, ok := sockets.Load(id)
	return &conn
}

func (manager *SocketManager) Ping(id string) {
	conn := manager.Get(id)
	if conn != nil {
		manager.writeText(*conn, "ping", "")
	}
}

func (manager *SocketManager) writeText(socket SocketConnection, event string, message string) {
	if socket.Writer == nil {
		return
	}
	var err error
	if event != "" {
		_, err = fmt.Fprintf(socket.Writer, "event: %s\ndata: %s\n\n", event, message)
	} else {
		_, err = fmt.Fprintf(socket.Writer, "data: %s\n\n", message)
	}
	if err != nil {
		manager.CloseWithError(socket.Id, 1008, "failed to write message")
	}
	socket.Flush <- true
}

func (manager *SocketManager) BroadcastText(roomId string, message string, predicate func(conn SocketConnection) bool) {
	sockets, ok := manager.sockets.Load(roomId)

	if !ok {
		return
	}

	sockets.Range(func(id string, conn SocketConnection) bool {
		if predicate(conn) {
			manager.writeText(conn, "", message)
		}
		return true
	})
}

func (manager *SocketManager) SendText(id string, message string) {
	conn := manager.Get(id)
	if conn != nil {
		manager.writeText(*conn, "", message)
	}
}
