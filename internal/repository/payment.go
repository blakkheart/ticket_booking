package repository

import (
	"context"
	"ticket-booking/internal/payment/models"
)

type PaymentRepository interface {
	Create(ctx context.Context, p *models.Payment) error
	Get(ctx context.Context, id string) (*models.Payment, error)
	Update(ctx context.Context, p *models.Payment) error
	GetByIntentID(ctx context.Context, intentID int64) ([]*models.Payment, error)
}
