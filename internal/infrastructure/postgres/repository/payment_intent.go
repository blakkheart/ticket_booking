package repository

import (
	"context"
	"log/slog"
	sqlc_repository "ticket-booking/internal/db/sqlc"
	"ticket-booking/internal/payment_intent/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type paymentIntentRepository struct {
	queries *sqlc_repository.Queries
	logger  *slog.Logger
}

func NewPaymentIntentRepository(
	db *pgxpool.Pool,
	logger *slog.Logger,
) *paymentIntentRepository {
	return &paymentIntentRepository{
		queries: sqlc_repository.New(db),
		logger:  logger,
	}
}

func NewPaymentIntentRepositoryFromQueries(
	q *sqlc_repository.Queries,
	logger *slog.Logger,
) *paymentIntentRepository {
	return &paymentIntentRepository{
		queries: q,
		logger:  logger,
	}
}

func (repo *paymentIntentRepository) Create(ctx context.Context, p *models.PaymentIntent) error{
	return nil
}

func (repo *paymentIntentRepository) Get(ctx context.Context, id int64) (*models.PaymentIntent, error) {
	return nil, nil
}

func (repo *paymentIntentRepository) GetByBookingID(ctx context.Context, bookingID int64) (*models.PaymentIntent, error) {
	return nil, nil
}

func (repo *paymentIntentRepository) GetMany(ctx context.Context, filter any) ([]*models.PaymentIntent, error) {
	return nil, nil
}

func (repo *paymentIntentRepository) Update(ctx context.Context, booking *models.PaymentIntent) error {
	return nil
}

func (repo *paymentIntentRepository) fromSqlcPaymentIntent(b *sqlc_repository.Booking) *models.PaymentIntent {
	return nil
}
