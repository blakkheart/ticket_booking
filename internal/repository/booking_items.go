package repository

import (
	"context"
	"ticket-booking/internal/booking_items/models"

	"github.com/google/uuid"
)

type BookingItemsRepository interface {
	Create(
		ctx context.Context,
		dto *models.BookingItemIn,
	) (*models.BookingItem, error)
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (*models.BookingItem, error)
	GetMany(filter any) ([]*models.BookingItem, error)
}
