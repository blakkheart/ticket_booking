package ticket

import (
	"context"
	"errors"
	"log/slog"
	"ticket-booking/internal/repository"
	"ticket-booking/internal/ticket/models"
	"ticket-booking/internal/uow"

	"github.com/google/uuid"
)

type Service interface {
	Get(id uuid.UUID) (*models.TicketType, error)
	GetMany(filters any) []*models.TicketType
	Create(ctx context.Context, t *models.TicketTypeIn) (*models.TicketType, error)
	Reserve(ctx context.Context,
		uow uow.UnitOfWork,
		ticketTypeID uuid.UUID,
		quantity int32,
		eventID uuid.UUID,
	) (*models.TicketType, error)
}

func NewService(repo repository.TicketRepository, logger *slog.Logger) Service {
	return &ticketService{Repo: repo, logger: logger}
}

type ticketService struct {
	Repo   repository.TicketRepository
	logger *slog.Logger
}

func (service *ticketService) Get(id uuid.UUID) (*models.TicketType, error) {
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
	ticketTypeID uuid.UUID,
	quantity int32,
	eventID uuid.UUID,
) (*models.TicketType, error) {

	repo := uow.TicketRepo()
	ticketType, err := repo.Get(ctx, ticketTypeID)

	if err != nil {
		service.logger.Error("ticket type doesn't exist", "error", err)
		return nil, errors.New("TicketType dosent exist")
	}
	if ticketType.EventID != eventID {
		return nil, errors.New("TicketType dosent match eventID")
	}

	if ticketType.AvailableQuantity < quantity {
		return nil, errors.New("Avaliable ticket quantity is too low")
	}

	ticket, err := repo.UpdateTicketQuantityByID(ctx, ticketTypeID, ticketType.AvailableQuantity-quantity)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}
