package tasks

import (
	"context"
	"github.com/google/uuid"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"time"
	"todolist/ent"
	"todolist/ent/predicate"
	"todolist/ent/task"
	"todolist/internal/util"
)

type Service struct {
	db        *ent.Client
	ipAddress string
}

type CreateRequest struct {
	Name string
	Tags []string
}

func NewService(ctx *h.RequestContext) Service {
	return Service{
		ipAddress: util.GetClientIp(ctx.Request),
		db:        service.Get[ent.Client](ctx.ServiceLocator()),
	}
}

func (s *Service) Create(request CreateRequest) (*ent.Task, error) {
	return s.db.Task.Create().
		SetName(request.Name).
		SetTags(request.Tags).
		SetIPAddress(s.ipAddress).
		Save(context.Background())
}

func (s *Service) Get(id uuid.UUID) (*ent.Task, error) {
	return s.db.Task.Get(context.Background(), id)
}

func (s *Service) SetName(id uuid.UUID, name string) (*ent.Task, error) {
	return s.db.Task.UpdateOneID(id).Where(task.IPAddress(s.ipAddress)).SetName(name).Save(context.Background())
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
		Where(task.IPAddress(s.ipAddress)).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	return err
}

func (s *Service) ClearCompleted() error {
	ctx := context.Background()
	_, err := s.db.Task.Delete().Where(task.CompletedAtNotNil(), task.IPAddress(s.ipAddress)).Exec(ctx)
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
		Where(task.IPAddress(s.ipAddress)).
		Save(ctx)
}

func (s *Service) List(ps ...predicate.Task) ([]*ent.Task, error) {
	ps = append(ps, task.IPAddress(s.ipAddress))
	return s.db.Task.Query().Where(ps...).All(context.Background())
}
