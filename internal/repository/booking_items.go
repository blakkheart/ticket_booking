package repository

import (
	"context"
	"ticket-booking/internal/booking_items/models"
)

type BookingItemsRepository interface {
	Create(ctx context.Context, dto *models.BookingItemIn) (*models.BookingItem, error)
	Delete(id int64) error
	Get(id int64) (*models.BookingItem, error)
	GetMany(filter any) ([]*models.BookingItem, error)
}
