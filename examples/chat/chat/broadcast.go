package chat

import (
	"chat/internal/db"
	"chat/ws"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
)

type Manager struct {
	socketManager *ws.SocketManager
	queries       *db.Queries
}

func NewManager(loader *service.Locator) *Manager {
	return &Manager{
		socketManager: service.Get[ws.SocketManager](loader),
		queries:       service.Get[db.Queries](loader),
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
				m.backFill(event.Id)
			case ws.DisconnectedEvent:
				fmt.Printf("User %s disconnected\n", event.Id)
			case ws.MessageEvent:
				m.onMessage(event.Id, event.Payload)
			}
		}
	}
}

func (m *Manager) backFill(socketId string) {
	messages, _ := m.queries.GetLastMessages(context.Background(), db.GetLastMessagesParams{
		ChatRoomID: "4ccc3f90a27c9375c98477571034b2e1",
		Limit:      50,
	})
	for _, message := range messages {
		m.socketManager.SendText(socketId,
			h.Render(MessageRow(message.Message)),
		)
	}
}

func (m *Manager) onMessage(socketId string, payload map[string]any) {
	fmt.Printf("Received message from %s: %v\n", socketId, payload)
	message := payload["message"].(string)

	if message == "" {
		return
	}

	ctx := context.Background()

	user, err := m.queries.CreateUser(ctx, uuid.NewString())

	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return
	}
	//chat, _ := m.queries.CreateChatRoom(ctx, "General")

	err = m.queries.InsertMessage(
		context.Background(),
		db.InsertMessageParams{
			ChatRoomID: "4ccc3f90a27c9375c98477571034b2e1",
			UserID:     user.ID,
			Message:    message,
		},
	)

	if err != nil {
		fmt.Printf("Error inserting message: %v\n", err)
		return
	}

	m.socketManager.BroadcastText(
		h.Render(MessageRow(message)),
	)
}
