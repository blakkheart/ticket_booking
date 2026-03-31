package ticket

import (
	"context"
	"log/slog"
)

type Service interface {
	Get(id int64) *TicketType
	GetMany(filters any) []*TicketType
	Create(ctx context.Context, t *TicketTypeIn) *TicketType
}

func NewService(repo Repository, logger *slog.Logger) Service {
	return &ticketService{Repo: repo, logger: logger}
}

type ticketService struct {
	Repo   Repository
	logger *slog.Logger
}

func (service *ticketService) Get(id int64) *TicketType {
	return nil
}

func (service *ticketService) GetMany(filters any) []*TicketType {
	return nil
}

func (service *ticketService) Create(ctx context.Context, t *TicketTypeIn) *TicketType {
	return nil
}
