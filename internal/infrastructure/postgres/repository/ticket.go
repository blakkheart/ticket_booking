package repository

import (
	"context"
	"log/slog"
	sqlc_repository "ticket-booking/internal/db/sqlc"
	"ticket-booking/internal/money"
	"ticket-booking/internal/ticket/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ticketRepository struct {
	queries *sqlc_repository.Queries
	logger  *slog.Logger
}

func NewTicketRepository(
	db *pgxpool.Pool,
	logger *slog.Logger,
) *ticketRepository {
	return &ticketRepository{
		queries: sqlc_repository.New(db),
		logger:  logger,
	}
}

func NewTicketRepositoryFromQueries(
	q *sqlc_repository.Queries,
	logger *slog.Logger,
) *ticketRepository {
	return &ticketRepository{
		queries: q,
		logger:  logger,
	}
}

func (repo *ticketRepository) Create(
	ctx context.Context,
	t *models.TicketTypeIn,
) (*models.TicketType, error) {
	var price pgtype.Numeric
	if err := price.Scan(t.Price.String()); err != nil {
		return nil, err
	}

	ticket, err := repo.queries.CreateTicket(
		ctx,
		sqlc_repository.CreateTicketParams{
			ID:                t.ID,
			Name:              t.Name,
			Description:       t.Description,
			Price:             price,
			AvailableQuantity: t.AvailableQuantity,
			EventID:           t.EventID,
		},
	)

	return repo.fromSqlcTicket(&ticket), err
}

func (repo *ticketRepository) Delete(id uuid.UUID) error {
	return nil
}

func (repo *ticketRepository) UpdateTicketQuantityByID(
	ctx context.Context,
	id uuid.UUID,
	newQuantity int32,
) (*models.TicketType, error) {
	ticket, err := repo.queries.UpdateTicketQuantityByID(
		ctx,
		sqlc_repository.UpdateTicketQuantityByIDParams{
			ID:                id,
			AvailableQuantity: newQuantity,
		},
	)
	if err != nil {
		return nil, err
	}
	return repo.fromSqlcTicket(&ticket), nil
}

func (repo *ticketRepository) Get(
	ctx context.Context,
	id uuid.UUID,
) (*models.TicketType, error) {
	ticketWithEvent, err := repo.queries.GetTicketByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return repo.fromSqlcTicket(&ticketWithEvent.TicketType), nil
}

func (repo *ticketRepository) GetMany(filter any) ([]*models.TicketType, error) {
	return nil, nil
}

func (repo *ticketRepository) fromSqlcTicket(
	t *sqlc_repository.TicketType,
) *models.TicketType {
	var priceStr string
	_ = t.Price.Scan(&priceStr)
	moneyType, _ := money.NewMoney(priceStr)

	ticket := &models.TicketType{
		ID:                t.ID,
		Name:              t.Name,
		Description:       t.Description,
		Price:             *moneyType,
		AvailableQuantity: t.AvailableQuantity,
		EventID:           t.EventID,
	}
	return ticket
}
