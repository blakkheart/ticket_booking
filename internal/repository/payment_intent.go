package repository

import (
	"context"
	"ticket-booking/internal/payment_intent/models"
)

type PaymentIntentRepository interface {
	Create(ctx context.Context, intent *models.PaymentIntent) error
	Get(ctx context.Context, id int64) (*models.PaymentIntent, error)
	GetByBookingID(ctx context.Context, bookingID int64) (*models.PaymentIntent, error)
	Update(ctx context.Context, intent *models.PaymentIntent) error
}
