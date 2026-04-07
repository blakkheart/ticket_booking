package repository

import (
	"context"
	"ticket-booking/internal/event/models"
)

type EventRepository interface {
	Create(ctx context.Context, dto *models.EventIn) (*models.Event, error)
	Delete(id int64) error
	Get(id int64) (*models.Event, error)
	GetMany(filter any) ([]*models.Event, error)
}
