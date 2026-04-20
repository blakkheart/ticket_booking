package event

import (
	"context"
	"log/slog"
	"ticket-booking/internal/event/models"
	"ticket-booking/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	Get(id uuid.UUID) (*models.Event, error)
	GetMany(filters any) []*models.Event
	Create(ctx context.Context, e *models.EventIn) (*models.Event, error)
}

func NewService(repo repository.EventRepository, logger *slog.Logger) Service {
	return &eventService{Repo: repo, logger: logger}
}

type eventService struct {
	Repo   repository.EventRepository
	logger *slog.Logger
}

func (service *eventService) Get(id uuid.UUID) (*models.Event, error) {
	event, err := service.Repo.Get(id)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (service *eventService) GetMany(filters any) []*models.Event {
	return nil
}

func (service *eventService) Create(ctx context.Context, e *models.EventIn) (*models.Event, error) {

	u := &models.EventIn{
		Title: "1",
	}

	event, err := service.Repo.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return event, nil
}
