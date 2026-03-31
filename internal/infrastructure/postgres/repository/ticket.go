package repository

import (
	"context"
	"log/slog"
	sqlc_repository "ticket-booking/internal/db/sqlc"
	"ticket-booking/internal/money"
	"ticket-booking/internal/ticket"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ticketRepository struct {
	queries *sqlc_repository.Queries
	logger  *slog.Logger
}

func NewTicketRepository(db *pgxpool.Pool, logger *slog.Logger) *ticketRepository {
	return &ticketRepository{
		queries: sqlc_repository.New(db),
		logger:  logger,
	}
}

func (repo *ticketRepository) Create(ctx context.Context, t *ticket.TicketTypeIn) (*ticket.TicketType, error) {
	var price pgtype.Numeric
	if err := price.Scan(t.Price.String()); err != nil {
		return nil, err
	}

	ticket, err := repo.queries.CreateTicket(
		ctx,
		sqlc_repository.CreateTicketParams{
			Name:              t.Name,
			Description:       t.Description,
			Price:             price,
			AvailableQuantity: t.AvailableQuantity,
			EventID:           t.EventID,
		},
	)

	return repo.fromSqlcAccount(&ticket), err
}

func (repo *ticketRepository) Delete(id int64) error {
	return nil
}

func (repo *ticketRepository) Get(id int64) (*ticket.TicketType, error) {
	return nil, nil
}

func (repo *ticketRepository) GetMany(filter any) ([]*ticket.TicketType, error) {
	return nil, nil
}

func (repo *ticketRepository) fromSqlcAccount(t *sqlc_repository.TicketType) *ticket.TicketType {
	var priceStr string
	_ = t.Price.Scan(&priceStr)
	moneyType, _ := money.NewMoney(priceStr)

	ticket := &ticket.TicketType{
		ID:                t.ID,
		Name:              t.Name,
		Description:       t.Description,
		Price:             *moneyType,
		AvailableQuantity: t.AvailableQuantity,
		EventID:           t.EventID,
	}
	return ticket
}
