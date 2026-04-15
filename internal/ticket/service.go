package ticket

import (
	"context"
	"log/slog"
	"ticket-booking/internal/repository"
	"ticket-booking/internal/ticket/models"
	"ticket-booking/internal/uow"
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

func (service *ticketService) Reserve(ctx context.Context,
	uow uow.UnitOfWork,
	ticketTypeID int64,
	qty int,
) (*models.TicketType, error) {

	// repo := uow.TicketRepo()

	// ticket, err := repo.Reserve(ctx, ticketTypeID, qty)
	// if err != nil {
	// 	return nil, err
	// }

	// return ticket, nil
	return nil, nil
}
