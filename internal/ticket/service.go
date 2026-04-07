package ticket

import (
	"context"
	"log/slog"
	"ticket-booking/internal/repository"
	"ticket-booking/internal/ticket/models"
)

type Service interface {
	Get(id int64) (*models.TicketType, error)
	GetMany(filters any) []*models.TicketType
	Create(ctx context.Context, t *models.TicketTypeIn) (*models.TicketType, error)
}

func NewService(repo repository.TicketRepository, logger *slog.Logger) Service {
	return &ticketService{Repo: repo, logger: logger}
}

type ticketService struct {
	Repo   repository.TicketRepository
	logger *slog.Logger
}

func (service *ticketService) Get(id int64) (*models.TicketType, error) {
	ticket, err := service.Get(id)
	if err != nil {
		return nil, err
	}
	return ticket, nil

}

func (service *ticketService) GetMany(filters any) []*models.TicketType {
	return nil
}

func (service *ticketService) Create(ctx context.Context, t *models.TicketTypeIn) (*models.TicketType, error) {
	ticket, err := service.Create(ctx, t)
	if err != nil {
		return nil, err
	}
	return ticket, nil

}
