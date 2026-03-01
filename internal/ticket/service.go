package ticket

import (
	"context"
)

type Service interface {
	Get(id int64) *Ticket
	GetMany(filters any) []*Ticket
	Create(ctx context.Context, user *Ticket) *Ticket
}

func NewService(repo Repository) Service {
	return &ticketService{Repo: repo}
}

type ticketService struct {
	Repo Repository
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
