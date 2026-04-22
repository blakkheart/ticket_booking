package repository

import (
	"context"
	"ticket-booking/internal/booking/models"

	"github.com/google/uuid"
)

type BookingRepository interface {
	Create(
		ctx context.Context,
		dto *models.BookingIn,
	) (*models.Booking, error)
	Delete(
		ctx context.Context,
		id uuid.UUID,
	) error
	Get(
		ctx context.Context,
		id uuid.UUID,
	) (*models.Booking, error)
	GetMany(
		ctx context.Context,
		filter any,
	) ([]*models.Booking, error)
	Update(
		ctx context.Context,
		booking *models.Booking,
	) error
}
