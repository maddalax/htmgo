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
	c := make(chan ws.SocketEvent, 1)
	m.socketManager.Listen(c)

	for {
		select {
		case event := <-c:
			switch event.Type {
			case ws.ConnectedEvent:
				go m.OnConnected(event)
			case ws.DisconnectedEvent:
				go m.OnDisconnected(event)
			case ws.MessageEvent:
				go m.onMessage(event)
			default:
				fmt.Printf("Unknown event type: %s\n", event.Type)
			}
		}
	}
}

func (m *Manager) OnConnected(e ws.SocketEvent) {
	room, _ := m.service.GetRoom(e.RoomId)

	if room == nil {
		m.socketManager.CloseWithMessage(e.Id, "invalid room")
		return
	}

	user, err := m.queries.GetUserBySessionId(context.Background(), e.Id)

	if err != nil {
		m.socketManager.CloseWithMessage(e.Id, "invalid user")
		return
	}

	fmt.Printf("User %s connected to %s\n", user.Name, e.RoomId)

	// backfill all existing clients to the connected client
	m.socketManager.ForEachSocket(e.RoomId, func(conn ws.SocketConnection) {
		user, err := m.queries.GetUserBySessionId(context.Background(), conn.Id)
		if err != nil {
			return
		}
		isMe := conn.Id == e.Id
		fmt.Printf("Sending connected user %s to %s\n", user.Name, e.Id)
		m.socketManager.SendText(e.Id, h.Render(ConnectedUsers(user.Name, isMe)))
	})

	// send the connected user to all existing clients
	m.socketManager.BroadcastText(
		e.RoomId,
		h.Render(ConnectedUsers(user.Name, false)),
		func(conn ws.SocketConnection) bool {
			return conn.Id != e.Id
		},
	)

	m.backFill(e.Id, e.RoomId)
}

func (m *Manager) OnDisconnected(e ws.SocketEvent) {
	user, err := m.queries.GetUserBySessionId(context.Background(), e.Id)
	if err != nil {
		return
	}
	room, err := m.service.GetRoom(e.RoomId)
	if err != nil {
		return
	}
	fmt.Printf("User %s disconnected from %s\n", user.Name, room.ID)
	m.socketManager.BroadcastText(room.ID, h.Render(ConnectedUser(user.Name, true, false)), func(conn ws.SocketConnection) bool {
		return conn.Id != e.Id
	})
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
			e.RoomId,
			h.Render(MessageRow(saved)),
			func(conn ws.SocketConnection) bool {
				return true
			},
		)
	}
}
