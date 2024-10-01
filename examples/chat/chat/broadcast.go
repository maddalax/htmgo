package chat

import (
	"chat/internal/db"
	"chat/ws"
	"context"
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"time"
)

type Manager struct {
	socketManager *ws.SocketManager
	queries       *db.Queries
	service       *Service
}

func NewManager(locator *service.Locator) *Manager {
	return &Manager{
		socketManager: service.Get[ws.SocketManager](locator),
		queries:       service.Get[db.Queries](locator),
		service:       NewService(locator),
	}
}

func (m *Manager) StartListener() {
	c := make(chan ws.SocketEvent)
	m.socketManager.Listen(c)

	for {
		select {
		case event := <-c:
			switch event.Type {
			case ws.ConnectedEvent:
				m.OnConnected(event)
			case ws.DisconnectedEvent:
				m.OnDisconnected(event)
			case ws.MessageEvent:
				m.onMessage(event)
			}
		}
	}
}

func (m *Manager) OnConnected(e ws.SocketEvent) {
	fmt.Printf("User %s connected to room %s\n", e.Id, e.RoomId)
	user, err := m.queries.GetUserBySessionId(context.Background(), e.Id)

	if err != nil {
		return
	}

	m.socketManager.BroadcastText(h.Render(ConnectedUsers(user.Name)))
	m.socketManager.ForEachSocket(e.RoomId, func(conn ws.SocketConnection) {
		if conn.Id == e.Id {
			return
		}
		user, err := m.queries.GetUserBySessionId(context.Background(), conn.Id)
		if err != nil {
			return
		}
		m.socketManager.SendText(e.Id, h.Render(ConnectedUsers(user.Name)))
	})

	go m.backFill(e.Id, e.RoomId)
}

func (m *Manager) OnDisconnected(e ws.SocketEvent) {
	fmt.Printf("User %s disconnected\n", e.Id)
	user, err := m.queries.GetUserBySessionId(context.Background(), e.Id)
	if err != nil {
		return
	}
	m.socketManager.BroadcastText(h.Render(ConnectedUser(user.Name, true)))
}

func (m *Manager) backFill(socketId string, roomId string) {
	messages, _ := m.queries.GetLastMessages(context.Background(), db.GetLastMessagesParams{
		ChatRoomID: roomId,
		Limit:      200,
	})
	for _, message := range messages {
		parsed, _ := time.Parse("2006-01-02 15:04:05", message.CreatedAt)
		m.socketManager.SendText(socketId,
			h.Render(MessageRow(&Message{
				UserId:    message.UserID,
				UserName:  message.UserName,
				Message:   message.Message,
				CreatedAt: parsed,
			})),
		)
	}
}

func (m *Manager) onMessage(e ws.SocketEvent) {
	fmt.Printf("Received message from %s: %v\n", e.Id, e.Payload)
	message := e.Payload["message"].(string)

	if message == "" {
		return
	}

	user, err := m.queries.GetUserBySessionId(context.Background(), e.Id)

	if err != nil {
		fmt.Printf("Error getting user: %v\n", err)
		return
	}

	saved := m.service.InsertMessage(
		&user,
		e.RoomId,
		message,
	)

	if saved != nil {
		m.socketManager.BroadcastText(
			h.Render(MessageRow(saved)),
		)
	}
}
