package event

import (
	"context"
	"log/slog"
	"ticket-booking/internal/event/models"
	"ticket-booking/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	Get(
		ctx context.Context,
		id uuid.UUID,
	) (*models.Event, error)
	GetMany(
		ctx context.Context,
		filters *models.EventFilter,
	) []*models.Event
	Create(
		ctx context.Context,
		e *models.EventIn,
	) (*models.Event, error)
	UpdateById(
		ctx context.Context,
		id uuid.UUID,
		e *models.EventUpdate,
	) (*models.Event, error)
}

func NewService(
	repo repository.EventRepository,
	logger *slog.Logger,
) Service {
	return &eventService{Repo: repo, logger: logger}
}

type eventService struct {
	Repo   repository.EventRepository
	logger *slog.Logger
}

func (service *eventService) Get(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	event, err := service.Repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (service *eventService) GetMany(ctx context.Context, filters *models.EventFilter) []*models.Event {
	events, err := service.Repo.GetMany(ctx, filters)
	if err != nil {
		return nil
	}
	return events
}

func (service *eventService) Create(
	ctx context.Context,
	e *models.EventIn,
) (*models.Event, error) {

	event, err := service.Repo.Create(ctx, e)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (service *eventService) UpdateById(
	ctx context.Context,
	id uuid.UUID,
	e *models.EventUpdate,
) (*models.Event, error) {
	event, err := service.Repo.UpdateById(ctx, id, e)
	if err != nil {
		return nil, err
	}
	return event, nil
}
