package repository

import (
	"context"
	"ticket-booking/internal/payment/models"
)

type PaymentRepository interface {
	Create(ctx context.Context, p *models.Payment) error
	GetByID(ctx context.Context, id string) (*models.Payment, error)
	Update(ctx context.Context, p *models.Payment) error
}
