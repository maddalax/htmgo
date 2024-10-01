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
				fmt.Printf("User %s connected\n", event.Id)
				m.backFill(event.Id, event.RoomId)
			case ws.DisconnectedEvent:
				fmt.Printf("User %s disconnected\n", event.Id)
			case ws.MessageEvent:
				m.onMessage(event)
			}
		}
	}
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
