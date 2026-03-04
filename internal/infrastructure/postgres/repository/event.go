package repository

import (
	"context"
	"log/slog"
	sqlc_repository "ticket-booking/internal/db/sqlc"
	eventModel "ticket-booking/internal/event"

	"github.com/jackc/pgx/v5/pgxpool"
)

type eventRepository struct {
	queries *sqlc_repository.Queries
	logger  *slog.Logger
}

func NewEventRepository(db *pgxpool.Pool, logger *slog.Logger) *eventRepository {
	return &eventRepository{
		queries: sqlc_repository.New(db),
		logger:  logger,
	}
}

func (repo *eventRepository) Create(ctx context.Context, user *eventModel.Event) (*eventModel.Event, error) {
	account, err := repo.queries.CreateAccount(
		ctx,
		sqlc_repository.CreateAccountParams{
			Name:     "1",
			Email:    "1",
			Password: "1",
		},
	)

	return repo.fromSqlcAccount(&account), err
}

func (repo *eventRepository) Delete(id int64) error {
	return nil
}

func (repo *eventRepository) Get(id int64) (*eventModel.Event, error) {
	return nil, nil
}

func (repo *eventRepository) GetMany(filter any) ([]*eventModel.Event, error) {
	return nil, nil
}

func (repo *eventRepository) fromSqlcAccount(a *sqlc_repository.Account) *eventModel.Event {
	account := &eventModel.Event{
		ID: a.ID,
	}
	return account
}
