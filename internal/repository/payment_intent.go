package repository

import (
	"context"
	"ticket-booking/internal/payment_intent/models"

	"github.com/google/uuid"
)

type PaymentIntentRepository interface {
	Create(ctx context.Context, intent *models.PaymentIntent) error
	Get(ctx context.Context, id uuid.UUID) (*models.PaymentIntent, error)
	GetByBookingID(ctx context.Context, bookingID uuid.UUID) (*models.PaymentIntent, error)
	Update(ctx context.Context, intent *models.PaymentIntent) error
}
