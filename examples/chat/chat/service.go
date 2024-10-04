package chat

import (
	"chat/internal/db"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/maddalax/htmgo/framework/service"
	"log"
	"time"
)

type Message struct {
	UserId    int64     `json:"userId"`
	UserName  string    `json:"userName"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

type Service struct {
	queries *db.Queries
}

func NewService(locator *service.Locator) *Service {
	return &Service{
		queries: service.Get[db.Queries](locator),
	}
}

func (s *Service) InsertMessage(user *db.User, roomId string, message string) *Message {
	err := s.queries.InsertMessage(context.Background(), db.InsertMessageParams{
		UserID:     user.ID,
		Username:   user.Name,
		ChatRoomID: roomId,
		Message:    message,
	})
	if err != nil {
		log.Printf("Failed to insert message: %v\n", err)
		return nil
	}
	return &Message{
		UserId:    user.ID,
		UserName:  user.Name,
		Message:   message,
		CreatedAt: time.Now(),
	}
}

func (s *Service) GetUserBySession(sessionId string) (*db.User, error) {
	user, err := s.queries.GetUserBySessionId(context.Background(), sessionId)
	return &user, err
}

func (s *Service) CreateUser(name string) (*db.CreateUserRow, error) {
	nameWithHash := fmt.Sprintf("%s#%s", name, uuid.NewString()[0:4])
	sessionId := fmt.Sprintf("session-%s-%s", uuid.NewString(), uuid.NewString())
	user, err := s.queries.CreateUser(context.Background(), db.CreateUserParams{
		Name:      nameWithHash,
		SessionID: sessionId,
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Service) CreateRoom(name string) (*db.CreateChatRoomRow, error) {
	room, err := s.queries.CreateChatRoom(context.Background(), db.CreateChatRoomParams{
		ID:   fmt.Sprintf("room-%s-%s", uuid.NewString()[0:8], name),
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (s *Service) GetRoom(id string) (*db.ChatRoom, error) {
	room, err := s.queries.GetChatRoom(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &room, nil
}
