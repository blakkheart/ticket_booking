package repository

import (
	"context"
	"ticket-booking/internal/payment/models"

	"github.com/google/uuid"
)

type PaymentRepository interface {
	Create(ctx context.Context, p *models.Payment) error
	Get(ctx context.Context, id string) (*models.Payment, error)
	Update(ctx context.Context, p *models.Payment) error
	GetByIntentID(ctx context.Context, intentID uuid.UUID) ([]*models.Payment, error)
}
