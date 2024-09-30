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
	Type    EventType
	Payload map[string]any
}

type SocketManager struct {
	sockets   *xsync.MapOf[string, *websocket.Conn]
	listeners []chan SocketEvent
}

func NewSocketManager() *SocketManager {
	return &SocketManager{
		sockets: xsync.NewMapOf[string, *websocket.Conn](),
	}
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
	manager.dispatch(SocketEvent{
		Id:      id,
		Type:    MessageEvent,
		Payload: message,
	})
}

func (manager *SocketManager) Add(id string, conn *websocket.Conn) {
	manager.sockets.Store(id, conn)
	manager.dispatch(SocketEvent{
		Id:      id,
		Type:    ConnectedEvent,
		Payload: map[string]any{},
	})
}

func (manager *SocketManager) OnClose(id string) {
	manager.dispatch(SocketEvent{
		Id:      id,
		Type:    DisconnectedEvent,
		Payload: map[string]any{},
	})
	manager.sockets.Delete(id)
}

func (manager *SocketManager) CloseWithError(id string, message string) {
	conn := manager.Get(id)
	if conn != nil {
		defer manager.OnClose(id)
		conn.Close(websocket.StatusInternalError, message)
	}
}

func (manager *SocketManager) Disconnect(id string) {
	conn := manager.Get(id)
	if conn != nil {
		defer manager.OnClose(id)
		_ = conn.CloseNow()
	}
}

func (manager *SocketManager) Get(id string) *websocket.Conn {
	conn, _ := manager.sockets.Load(id)
	return conn
}

func (manager *SocketManager) Broadcast(message []byte, messageType websocket.MessageType) {
	ctx := context.Background()
	manager.sockets.Range(func(id string, conn *websocket.Conn) bool {
		err := conn.Write(ctx, messageType, message)
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
		_ = conn.Write(context.Background(), websocket.MessageText, []byte(message))
	}
}
