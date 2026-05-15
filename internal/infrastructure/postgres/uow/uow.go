package postgresuow

import (
	"context"
	"errors"
	"log/slog"
	sqlc_repository "ticket-booking/internal/db/sqlc"
	"ticket-booking/internal/infrastructure/postgres/repository"
	repositoryInteface "ticket-booking/internal/repository"

	"github.com/jackc/pgx/v5"
)

type UOW interface {
	bookingRepo() repositoryInteface.BookingRepository
	ticketRepo() repositoryInteface.TicketRepository
	bookingItemsRepo() repositoryInteface.BookingItemsRepository
	paymentIntentRepo() repositoryInteface.PaymentIntentRepository
	paymentRepo() repositoryInteface.PaymentRepository

	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type uow struct {
	tx      pgx.Tx
	queries *sqlc_repository.Queries
	logger *slog.Logger

	bookingRepo       repositoryInteface.BookingRepository
	eventRepo         repositoryInteface.EventRepository
	ticketRepo        repositoryInteface.TicketRepository
	bookingItemsRepo  repositoryInteface.BookingItemsRepository
	paymentIntentRepo repositoryInteface.PaymentIntentRepository
	paymentRepo       repositoryInteface.PaymentRepository
}

func newUow(tx pgx.Tx, logger *slog.Logger) *uow {
	q := sqlc_repository.New(tx)

	return &uow{
		tx:      tx,
		queries: q,
		logger: logger,

		bookingRepo:       repository.NewBookingRepositoryFromQueries(q, logger),
		eventRepo:         repository.NewEventRepositoryFromQueries(q, logger),
		ticketRepo:        repository.NewTicketRepositoryFromQueries(q, logger),
		bookingItemsRepo:  repository.NewBookingItemsRepositoryFromQueries(q, logger),
		paymentIntentRepo: repository.NewPaymentIntentRepositoryFromQueries(q, logger),
		paymentRepo:       repository.NewPaymentRepositoryFromQueries(q, logger),
	}
}

func (u *uow) BookingRepo() repositoryInteface.BookingRepository {
	return u.bookingRepo
}

func (u *uow) EventRepo() repositoryInteface.EventRepository {
	return u.eventRepo
}

func (u *uow) TicketRepo() repositoryInteface.TicketRepository {
	return u.ticketRepo
}

func (u *uow) BookingItemsRepo() repositoryInteface.BookingItemsRepository {
	return u.bookingItemsRepo
}
func (u *uow) PaymentRepo() repositoryInteface.PaymentRepository {
	return u.paymentRepo
}
func (u *uow) PaymentIntentRepo() repositoryInteface.PaymentIntentRepository {
	return u.paymentIntentRepo
}

func (u *uow) Commit(ctx context.Context) error {
	return u.tx.Commit(ctx)
}

func (u *uow) Rollback(ctx context.Context) error {
	err := u.tx.Rollback(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrTxClosed) {
			u.logger.Debug("rollback skipped. tx already closed")
			return nil
		}

		u.logger.Error("rollback failed", "error", err)
		return err
	}

	u.logger.Debug("transaction rolled back")
	return nil
}
