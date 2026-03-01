package postgres

import (
	"context"
	"ticket-booking/internal/booking"
	sqlc_repository "ticket-booking/repository/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type bookingRepository struct {
	queries *sqlc_repository.Queries
}

func NewBookingRepository(db *pgxpool.Pool) *bookingRepository {
	return &bookingRepository{
		queries: sqlc_repository.New(db),
	}
}

func (repo *bookingRepository) Create(ctx context.Context, user *booking.Booking) (*booking.Booking, error) {
	account, err := repo.queries.CreateAccount(
		ctx,
		sqlc_repository.CreateAccountParams{},
	)

	return repo.fromSqlcAccount(&account), err
}

func (repo *bookingRepository) Delete(id int64) error {
	return nil
}

func (repo *bookingRepository) Get(id int64) (*booking.Booking, error) {
	return nil, nil
}

func (repo *bookingRepository) GetMany(filter any) ([]*booking.Booking, error) {
	return nil, nil
}

func (repo *bookingRepository) fromSqlcAccount(a *sqlc_repository.Account) *booking.Booking {
	account := &booking.Booking{
		ID: a.ID,
	}
	return account
}
