package ws

import (
	"context"
	"github.com/coder/websocket"
	"github.com/puzpuzpuz/xsync/v3"
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

type SocketConnection struct {
	Id     string
	Conn   *websocket.Conn
	RoomId string
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

func (manager *SocketManager) Add(roomId string, id string, conn *websocket.Conn) {
	manager.idToRoom.Store(id, roomId)

	sockets, ok := manager.sockets.LoadOrCompute(roomId, func() *xsync.MapOf[string, SocketConnection] {
		return xsync.NewMapOf[string, SocketConnection]()
	})

	sockets.Store(id, SocketConnection{
		Id:     id,
		Conn:   conn,
		RoomId: roomId,
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

func (manager *SocketManager) CloseWithError(id string, code websocket.StatusCode, message string) {
	conn := manager.Get(id)
	if conn != nil {
		go manager.OnClose(id)
		conn.Conn.Close(code, message)
	}
}

func (manager *SocketManager) Disconnect(id string) {
	conn := manager.Get(id)
	if conn != nil {
		go manager.OnClose(id)
		_ = conn.Conn.CloseNow()
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

func (manager *SocketManager) Broadcast(roomId string, message []byte, messageType websocket.MessageType, predicate func(conn SocketConnection) bool) {
	ctx := context.Background()
	sockets, ok := manager.sockets.Load(roomId)

	if !ok {
		return
	}

	sockets.Range(func(id string, conn SocketConnection) bool {
		if predicate(conn) {
			conn.Conn.Write(ctx, messageType, message)
		}
		return true
	})
}

func (manager *SocketManager) BroadcastText(roomId string, message string, predicate func(conn SocketConnection) bool) {
	manager.Broadcast(roomId, []byte(message), websocket.MessageText, predicate)
}

func (manager *SocketManager) SendText(id string, message string) {
	conn := manager.Get(id)
	if conn != nil {
		_ = conn.Conn.Write(context.Background(), websocket.MessageText, []byte(message))
	}
}
