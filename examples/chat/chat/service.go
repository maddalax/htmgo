package chat

import (
	"chat/internal/db"
	"context"
	"github.com/maddalax/htmgo/framework/service"
)

type Service struct {
	queries *db.Queries
}

func NewService(locator *service.Locator) *Service {
	return &Service{
		queries: service.Get[db.Queries](locator),
	}
}

func (s *Service) GetRoom(id string) (*db.ChatRoom, error) {
	room, err := s.queries.GetChatRoom(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &room, nil
}
