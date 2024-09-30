package ws

import (
	"context"
	"fmt"
	"github.com/coder/websocket"
	"github.com/puzpuzpuz/xsync/v3"
)

type MessageEvent struct {
	Id      string
	Message map[string]any
}

type SocketManager struct {
	sockets   *xsync.MapOf[string, *websocket.Conn]
	listeners []chan MessageEvent
}

func NewSocketManager() *SocketManager {
	return &SocketManager{
		sockets: xsync.NewMapOf[string, *websocket.Conn](),
	}
}

func (manager *SocketManager) Listen(listener chan MessageEvent) {
	if manager.listeners == nil {
		manager.listeners = make([]chan MessageEvent, 0)
	}
	manager.listeners = append(manager.listeners, listener)
}

func (manager *SocketManager) OnMessage(id string, message map[string]any) {
	for _, listener := range manager.listeners {
		listener <- MessageEvent{
			Id:      id,
			Message: message,
		}
	}
}

func (manager *SocketManager) Add(id string, conn *websocket.Conn) {
	manager.sockets.Store(id, conn)
}

func (manager *SocketManager) CloseWithError(id string, message string) {
	conn := manager.Get(id)
	if conn != nil {
		conn.Close(websocket.StatusInternalError, message)
	}
}

func (manager *SocketManager) Disconnect(id string) {
	conn := manager.Get(id)
	if conn != nil {
		_ = conn.CloseNow()
	}
	manager.sockets.Delete(id)
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
