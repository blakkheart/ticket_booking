package repository

import (
	"context"
	"ticket-booking/internal/ticket/models"

	"github.com/google/uuid"
)

type TicketRepository interface {
	Create(ctx context.Context, dto *models.TicketTypeIn) (*models.TicketType, error)
	Delete(id uuid.UUID) error
	Get(ctx context.Context, id uuid.UUID) (*models.TicketType, error)
	GetMany(filter any) ([]*models.TicketType, error)
	UpdateTicketQuantityByID(ctx context.Context, id uuid.UUID, newQuantity int32) (*models.TicketType, error)
}
