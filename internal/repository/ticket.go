package repository

import (
	"context"
	"ticket-booking/internal/ticket/models"
)

type TicketRepository interface {
	Create(ctx context.Context, dto *models.TicketTypeIn) (*models.TicketType, error)
	Delete(id int64) error
	Get(ctx context.Context, id int64) (*models.TicketType, error)
	GetMany(filter any) ([]*models.TicketType, error)
	UpdateTicketQuantityByID(ctx context.Context, id int64, newQuantity int32) (*models.TicketType, error)
}
