package repository

import (
	"context"
	"log/slog"
	sqlc_repository "ticket-booking/internal/db/sqlc"
	"ticket-booking/internal/payment/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type paymentRepository struct {
	queries *sqlc_repository.Queries
	logger  *slog.Logger
}

func NewPaymentRepository(
	db *pgxpool.Pool,
	logger *slog.Logger,
) *paymentRepository {
	return &paymentRepository{
		queries: sqlc_repository.New(db),
		logger:  logger,
	}
}

func NewPaymentRepositoryFromQueries(
	q *sqlc_repository.Queries,
	logger *slog.Logger,
) *paymentRepository {
	return &paymentRepository{
		queries: q,
		logger:  logger,
	}
}

func (repo *paymentRepository) Create(ctx context.Context, p *models.Payment) error {
	return nil
}

func (repo *paymentRepository) Get(ctx context.Context, id string) (*models.Payment, error) {
	return nil, nil
}

func (repo *paymentRepository) GetByIntentID(ctx context.Context, intentID uuid.UUID) ([]*models.Payment, error) {
	return nil, nil
}

func (repo *paymentRepository) GetMany(ctx context.Context, filter any) ([]*models.Payment, error) {
	return nil, nil
}

func (repo *paymentRepository) Update(ctx context.Context, booking *models.Payment) error {
	return nil
}

func (repo *paymentRepository) fromSqlcPaymentIntent(b *sqlc_repository.Booking) *models.Payment {
	return nil
}
