package ws

import (
	"context"
	"fmt"
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
	sockets   *xsync.MapOf[string, SocketConnection]
	listeners []chan SocketEvent
}

func NewSocketManager() *SocketManager {
	return &SocketManager{
		sockets: xsync.NewMapOf[string, SocketConnection](),
	}
}

func (manager *SocketManager) ForEachSocket(roomId string, cb func(conn SocketConnection)) {
	manager.sockets.Range(func(id string, conn SocketConnection) bool {
		if conn.RoomId == roomId {
			cb(conn)
		}
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
	manager.sockets.Store(id, SocketConnection{
		Id:     id,
		Conn:   conn,
		RoomId: roomId,
	})
	s, ok := manager.sockets.Load(id)
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
	conn, ok := manager.sockets.Load(id)
	if !ok {
		return nil
	}
	return &conn
}

func (manager *SocketManager) Broadcast(message []byte, messageType websocket.MessageType) {
	ctx := context.Background()
	manager.sockets.Range(func(id string, conn SocketConnection) bool {
		err := conn.Conn.Write(ctx, messageType, message)
		if err != nil {
			manager.Disconnect(id)
		}
		return true
	})
}

func (manager *SocketManager) BroadcastText(message string) {
	fmt.Printf("Broadcasting message: \n%s\n", message)
	manager.Broadcast([]byte(message), websocket.MessageText)
}

func (manager *SocketManager) SendText(id string, message string) {
	conn := manager.Get(id)
	if conn != nil {
		_ = conn.Conn.Write(context.Background(), websocket.MessageText, []byte(message))
	}
}
