package wsutil

import (
	"fmt"
	"github.com/maddalax/htmgo/extensions/websocket/opts"
	"github.com/puzpuzpuz/xsync/v3"
	"strings"
	"time"
)

type EventType string
type WriterChan chan string
type DoneChan chan bool

const (
	ConnectedEvent    EventType = "connected"
	DisconnectedEvent EventType = "disconnected"
	MessageEvent      EventType = "message"
)

type SocketEvent struct {
	SessionId string
	RoomId    string
	Type      EventType
	Payload   map[string]any
}

type CloseEvent struct {
	Code   int
	Reason string
}

type SocketConnection struct {
	Id     string
	RoomId string
	Done   DoneChan
	Writer WriterChan
}

type SocketManager struct {
	sockets   *xsync.MapOf[string, *xsync.MapOf[string, SocketConnection]]
	idToRoom  *xsync.MapOf[string, string]
	listeners []chan SocketEvent
	opts      *opts.ExtensionOpts
}

func NewSocketManager(opts *opts.ExtensionOpts) *SocketManager {
	return &SocketManager{
		sockets:  xsync.NewMapOf[string, *xsync.MapOf[string, SocketConnection]](),
		idToRoom: xsync.NewMapOf[string, string](),
		opts:     opts,
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
	if listener != nil {
		manager.listeners = append(manager.listeners, listener)
	}
}

func (manager *SocketManager) dispatch(event SocketEvent) {
	fmt.Printf("dispatching event: %s\n", event.Type)
	done := make(chan struct{}, 1)
	go func() {
		for {
			select {
			case <-done:
				fmt.Printf("dispatched event: %s\n", event.Type)
				return
			case <-time.After(5 * time.Second):
				fmt.Printf("havent dispatched event after 5s, chan blocked: %s\n", event.Type)
			}
		}
	}()
	for _, listener := range manager.listeners {
		listener <- event
	}
	done <- struct{}{}
}

func (manager *SocketManager) OnMessage(id string, message map[string]any) {
	socket := manager.Get(id)
	if socket == nil {
		return
	}
	manager.dispatch(SocketEvent{
		SessionId: id,
		Type:      MessageEvent,
		Payload:   message,
		RoomId:    socket.RoomId,
	})
}

func (manager *SocketManager) Add(roomId string, id string, writer WriterChan, done DoneChan) {
	manager.idToRoom.Store(id, roomId)

	sockets, ok := manager.sockets.LoadOrCompute(roomId, func() *xsync.MapOf[string, SocketConnection] {
		return xsync.NewMapOf[string, SocketConnection]()
	})

	sockets.Store(id, SocketConnection{
		Id:     id,
		Writer: writer,
		RoomId: roomId,
		Done:   done,
	})

	s, ok := sockets.Load(id)
	if !ok {
		return
	}

	manager.dispatch(SocketEvent{
		SessionId: s.Id,
		Type:      ConnectedEvent,
		RoomId:    s.RoomId,
		Payload:   map[string]any{},
	})

	fmt.Printf("User %s connected to %s\n", id, roomId)
}

func (manager *SocketManager) OnClose(id string) {
	socket := manager.Get(id)
	if socket == nil {
		return
	}
	manager.dispatch(SocketEvent{
		SessionId: id,
		Type:      DisconnectedEvent,
		RoomId:    socket.RoomId,
		Payload:   map[string]any{},
	})
	manager.sockets.Delete(id)
}

func (manager *SocketManager) CloseWithMessage(id string, message string) {
	conn := manager.Get(id)
	if conn != nil {
		defer manager.OnClose(id)
		manager.writeText(*conn, message)
		conn.Done <- true
	}
}

func (manager *SocketManager) Disconnect(id string) {
	conn := manager.Get(id)
	if conn != nil {
		manager.OnClose(id)
		conn.Done <- true
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
		manager.writeText(*conn, "ping")
	}
}

func (manager *SocketManager) writeCloseRaw(writer WriterChan, message string) {
	manager.writeTextRaw(writer, message)
}

func (manager *SocketManager) writeTextRaw(writer WriterChan, message string) {
	timeout := 3 * time.Second
	select {
	case writer <- message:
	case <-time.After(timeout):
		fmt.Printf("could not send %s to channel after %s\n", message, timeout)
	}
}

func (manager *SocketManager) writeText(socket SocketConnection, message string) {
	if socket.Writer == nil {
		return
	}
	manager.writeTextRaw(socket.Writer, message)
}

func (manager *SocketManager) BroadcastText(roomId string, message string, predicate func(conn SocketConnection) bool) {
	sockets, ok := manager.sockets.Load(roomId)

	if !ok {
		return
	}

	sockets.Range(func(id string, conn SocketConnection) bool {
		if predicate(conn) {
			manager.writeText(conn, message)
		}
		return true
	})
}

func (manager *SocketManager) SendHtml(id string, message string) {
	conn := manager.Get(id)
	minified := strings.ReplaceAll(message, "\n", "")
	minified = strings.ReplaceAll(minified, "\t", "")
	minified = strings.TrimSpace(minified)
	if conn != nil {
		manager.writeText(*conn, minified)
	}
}

func (manager *SocketManager) SendText(id string, message string) {
	conn := manager.Get(id)
	if conn != nil {
		manager.writeText(*conn, message)
	}
}
