package event

import (
	"context"
	"log/slog"
)

type Service interface {
	Get(id int64) (*Event, error)
	GetMany(filters any) []*Event
	Create(ctx context.Context, e *EventIn) (*Event, error)
}

func NewService(repo Repository, logger *slog.Logger) Service {
	return &eventService{Repo: repo, logger: logger}
}

type eventService struct {
	Repo   Repository
	logger *slog.Logger
}

func (service *eventService) Get(id int64) (*Event, error) {
	event, err := service.Repo.Get(id)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (service *eventService) GetMany(filters any) []*Event {
	return nil
}

func (service *eventService) Create(ctx context.Context, e *EventIn) (*Event, error) {

	u := &EventIn{
		Title: "1",
	}

	event, err := service.Repo.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return event, nil
}
