package event

import (
	"context"
	"log"
	"log/slog"
)

type Service interface {
	Get(id int64) *Event
	GetMany(filters any) []*Event
	Create(ctx context.Context, user *Event) *Event
}

func NewService(repo Repository, logger *slog.Logger) Service {
	return &eventService{Repo: repo, logger: logger}
}

type eventService struct {
	Repo   Repository
	logger *slog.Logger
}

func (service *eventService) Get(id int64) *Event {
	value, err := service.Repo.Get(id)
	if err != nil {
		log.Fatal("Something wrong with repo")
	}
	return value
}

func (service *eventService) GetMany(filters any) []*Event {
	return nil
}

func (service *eventService) Create(ctx context.Context, user *Event) *Event {

	u := &Event{
		ID: 1,
	}

	value, err := service.Repo.Create(ctx, u)
	if err != nil {
		log.Fatal("Something wrong with repo")
	}

	return value
}
