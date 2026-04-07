package repository

import (
	"context"
	"ticket-booking/internal/booking/models"
)

type BookingRepository interface {
	Create(ctx context.Context, dto *models.BookingIn) (*models.Booking, error)
	Delete(id int64) error
	Get(id int64) (*models.Booking, error)
	GetMany(filter any) ([]*models.Booking, error)
}
