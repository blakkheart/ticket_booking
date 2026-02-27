package bookrep

import (
	"context"
	bookingModel "ticket-booking/models/domain/booking"
	sqlc_repository "ticket-booking/repository/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type bookingRepository struct {
	DB      *pgxpool.Pool
	queries *sqlc_repository.Queries
}

func New(db *pgxpool.Pool) *bookingRepository {
	return &bookingRepository{
		DB:      db,
		queries: sqlc_repository.New(db),
	}
}

func (repo *bookingRepository) Create(ctx context.Context, user *bookingModel.Booking) (*bookingModel.Booking, error) {
	account, err := repo.queries.CreateAccount(
		ctx,
		sqlc_repository.CreateAccountParams{},
	)

	return repo.fromSqlcAccount(&account), err
}

func (repo *bookingRepository) Delete(id int64) error {
	return nil
}

func (repo *bookingRepository) Get(id int64) (*bookingModel.Booking, error) {
	return nil, nil
}

func (repo *bookingRepository) GetMany(filter any) ([]*bookingModel.Booking, error) {
	return nil, nil
}

func (repo *bookingRepository) fromSqlcAccount(a *sqlc_repository.Account) *bookingModel.Booking {
	account := &bookingModel.Booking{
		ID: a.ID,
	}
	return account
}
