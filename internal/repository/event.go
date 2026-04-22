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
	Get(id uuid.UUID) (*models.Event, error)
	GetMany(filter any) ([]*models.Event, error)
}
