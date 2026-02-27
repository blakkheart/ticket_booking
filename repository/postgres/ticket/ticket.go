package ticketrep

import (
	"context"
	ticketModel "ticket-booking/models/domain/ticket"
	sqlc_repository "ticket-booking/repository/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ticketRepository struct {
	DB      *pgxpool.Pool
	queries *sqlc_repository.Queries
}

func New(db *pgxpool.Pool) *ticketRepository {
	return &ticketRepository{
		DB:      db,
		queries: sqlc_repository.New(db),
	}
}

func (repo *ticketRepository) Create(ctx context.Context, user *ticketModel.Ticket) (*ticketModel.Ticket, error) {
	account, err := repo.queries.CreateAccount(
		ctx,
		sqlc_repository.CreateAccountParams{},
	)

	return repo.fromSqlcAccount(&account), err
}

func (repo *ticketRepository) Delete(id int64) error {
	return nil
}

func (repo *ticketRepository) Get(id int64) (*ticketModel.Ticket, error) {
	return nil, nil
}

func (repo *ticketRepository) GetMany(filter any) ([]*ticketModel.Ticket, error) {
	return nil, nil
}

func (repo *ticketRepository) fromSqlcAccount(a *sqlc_repository.Account) *ticketModel.Ticket {
	account := &ticketModel.Ticket{
		ID: a.ID,
	}
	return account
}
