package repository

import (
	"context"
	"log/slog"
	"ticket-booking/internal/booking"
	sqlc_repository "ticket-booking/internal/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type bookingRepository struct {
	queries *sqlc_repository.Queries
	logger  *slog.Logger
}

func NewBookingRepository(db *pgxpool.Pool, logger *slog.Logger) *bookingRepository {
	return &bookingRepository{
		queries: sqlc_repository.New(db),
		logger:  logger,
	}
}

func (repo *bookingRepository) Create(ctx context.Context, b *booking.BookingIn) (*booking.Booking, error) {
	booking, err := repo.queries.CreateBooking(
		ctx,
		sqlc_repository.CreateBookingParams{
			AccountID: b.AccountID,
			Status:    b.Status,
			ExpiresAt: b.ExpiresAt,
			PaidAt:    b.PaidAt,
		},
	)

	return repo.fromSqlcBooking(&booking), err
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

func (repo *bookingRepository) fromSqlcBooking(b *sqlc_repository.Booking) *booking.Booking {
	booking := &booking.Booking{
		ID:        b.ID,
		AccountID: b.AccountID,
		Status:    b.Status,
		ExpiresAt: b.ExpiresAt,
		PaidAt:    b.PaidAt,
	}
	return booking
}
