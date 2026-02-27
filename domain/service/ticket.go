package service

import (
	"context"

	model_interface "ticket-booking/models/domain/interface"
	model "ticket-booking/models/domain/ticket"
)

type TicketService interface {
	Get(id int64) *model.Ticket
	GetMany(filters any) []*model.Ticket
	Create(ctx context.Context, user *model.Ticket) *model.Ticket
}

func NewT(repo model_interface.ITicketRepository) TicketService {
	return &ticketService{Repo: repo}
}

type ticketService struct {
	Repo model_interface.ITicketRepository
}

func (service *ticketService) Get(id int64) *model.Ticket {
	return nil
}

func (service *ticketService) GetMany(filters any) []*model.Ticket {
	return nil
}

func (service *ticketService) Create(ctx context.Context, user *model.Ticket) *model.Ticket {
	return nil
}
