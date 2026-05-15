package postgresuow

import (
	"context"
	"log/slog"
	uowmodel "ticket-booking/internal/uow"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type uowManager struct {
	pool   *pgxpool.Pool
	logger *slog.Logger
}

func NewManager(pool *pgxpool.Pool, logger *slog.Logger) *uowManager {
	return &uowManager{
		pool:   pool,
		logger: logger,
	}
}

func (m *uowManager) Begin(ctx context.Context) (uowmodel.UnitOfWork, error) {
	tx, err := m.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	return newUow(tx, m.logger), nil
}

func (m *uowManager) Do(
	ctx context.Context,
	fn func(uow uowmodel.UnitOfWork) error,
) error {
	uow, err := m.Begin(ctx)
	if err != nil {
		return err
	}

	defer uow.Rollback(ctx)

	if err := fn(uow); err != nil {
		return err
	}
	return uow.Commit(ctx)
}
