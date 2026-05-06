package repository

import (
	"context"
	"ticket-booking/internal/event/models"

	"github.com/google/uuid"
)

type EventRepository interface {
	Create(
		ctx context.Context,
		dto *models.EventIn,
	) (*models.Event, error)
	Delete(id uuid.UUID) error
	Get(
		ctx context.Context,
		id uuid.UUID,
	) (*models.Event, error)
	GetMany(
		ctx context.Context,
		filter *models.EventFilter,
	) ([]*models.Event, error)
	UpdateById(
		ctx context.Context,
		id uuid.UUID,
		e *models.EventUpdate,
	) (*models.Event, error)
}
