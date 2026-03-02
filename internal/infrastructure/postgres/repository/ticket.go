package repository

import (
	"context"
	"ticket-booking/internal/ticket"
	sqlc_repository "ticket-booking/internal/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ticketRepository struct {
	DB      *pgxpool.Pool
	queries *sqlc_repository.Queries
}

func NewTicketRepository(db *pgxpool.Pool) *ticketRepository {
	return &ticketRepository{
		DB:      db,
		queries: sqlc_repository.New(db),
	}
}

func (repo *ticketRepository) Create(ctx context.Context, user *ticket.Ticket) (*ticket.Ticket, error) {
	account, err := repo.queries.CreateAccount(
		ctx,
		sqlc_repository.CreateAccountParams{},
	)

	return repo.fromSqlcAccount(&account), err
}

func (repo *ticketRepository) Delete(id int64) error {
	return nil
}

func (repo *ticketRepository) Get(id int64) (*ticket.Ticket, error) {
	return nil, nil
}

func (repo *ticketRepository) GetMany(filter any) ([]*ticket.Ticket, error) {
	return nil, nil
}

func (repo *ticketRepository) fromSqlcAccount(a *sqlc_repository.Account) *ticket.Ticket {
	account := &ticket.Ticket{
		ID: a.ID,
	}
	return account
}
