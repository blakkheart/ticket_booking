package repository

import (
	"context"
	"log/slog"
	"ticket-booking/internal/booking/models"
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

func NewBookingRepositoryFromQueries(q *sqlc_repository.Queries, logger *slog.Logger) *bookingRepository {
	return &bookingRepository{
		queries: q,
		logger:  logger,
	}
}

func (repo *bookingRepository) Create(ctx context.Context, b *models.BookingIn) (*models.Booking, error) {
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

func (repo *bookingRepository) Get(id int64) (*models.Booking, error) {
	return nil, nil
}

func (repo *bookingRepository) GetMany(filter any) ([]*models.Booking, error) {
	return nil, nil
}

func (repo *bookingRepository) fromSqlcBooking(b *sqlc_repository.Booking) *models.Booking {
	booking := &models.Booking{
		ID:        b.ID,
		AccountID: b.AccountID,
		Status:    b.Status,
		ExpiresAt: b.ExpiresAt,
		PaidAt:    b.PaidAt,
	}
	return booking
}
