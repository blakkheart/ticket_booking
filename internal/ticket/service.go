package ticket

import (
	"context"
	"log/slog"
)

type Service interface {
	Get(id int64) *Ticket
	GetMany(filters any) []*Ticket
	Create(ctx context.Context, user *Ticket) *Ticket
}

func NewService(repo Repository, logger *slog.Logger) Service {
	return &ticketService{Repo: repo, logger: logger}
}

type ticketService struct {
	Repo   Repository
	logger *slog.Logger
}

func (service *ticketService) Get(id int64) *Ticket {
	return nil
}

func (service *ticketService) GetMany(filters any) []*Ticket {
	return nil
}

func (service *ticketService) Create(ctx context.Context, user *Ticket) *Ticket {
	return nil
}
