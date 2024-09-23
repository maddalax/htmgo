package tasks

import (
	"context"
	"github.com/google/uuid"
	"github.com/maddalax/htmgo/framework/service"
	"time"
	"todolist/ent"
	"todolist/ent/predicate"
	"todolist/ent/task"
)

type Service struct {
	db *ent.Client
}

type CreateRequest struct {
	Name string
	Tags []string
}

func NewService(locator *service.Locator) Service {
	return Service{
		db: service.Get[ent.Client](locator),
	}
}

func (s *Service) Create(request CreateRequest) (*ent.Task, error) {
	return s.db.Task.Create().
		SetName(request.Name).
		SetTags(request.Tags).
		Save(context.Background())
}

func (s *Service) Get(id uuid.UUID) (*ent.Task, error) {
	return s.db.Task.Get(context.Background(), id)
}

func (s *Service) SetName(id uuid.UUID, name string) (*ent.Task, error) {
	return s.db.Task.UpdateOneID(id).SetName(name).Save(context.Background())
}

func (s *Service) SetAllCompleted(value bool) error {
	ctx := context.Background()
	updater := s.db.Task.Update()

	if value {
		updater = updater.SetCompletedAt(time.Now())
	} else {
		updater = updater.ClearCompletedAt()
	}

	_, err := updater.
		SetUpdatedAt(time.Now()).
		Save(ctx)

	return err
}

func (s *Service) ClearCompleted() error {
	ctx := context.Background()
	_, err := s.db.Task.Delete().Where(task.CompletedAtNotNil()).Exec(ctx)
	return err
}

func (s *Service) SetCompleted(id uuid.UUID, value bool) (*ent.Task, error) {
	ctx := context.Background()
	updater := s.db.Task.UpdateOneID(id)

	if value {
		updater = updater.SetCompletedAt(time.Now())
	} else {
		updater = updater.ClearCompletedAt()
	}

	return updater.
		SetUpdatedAt(time.Now()).
		Save(ctx)
}

func (s *Service) List(ps ...predicate.Task) ([]*ent.Task, error) {
	return s.db.Task.Query().Where(ps...).All(context.Background())
}
