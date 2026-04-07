package postgresuow

import (
	"context"
	"log/slog"
	sqlc_repository "ticket-booking/internal/db/sqlc"
	"ticket-booking/internal/infrastructure/postgres/repository"
	repositoryInteface "ticket-booking/internal/repository"

	"github.com/jackc/pgx/v5"
)

type uow struct {
	tx      pgx.Tx
	queries *sqlc_repository.Queries

	bookingRepo      repositoryInteface.BookingRepository
	ticketRepo       repositoryInteface.TicketRepository
	bookingItemsRepo repositoryInteface.BookingItemsRepository
}

func newUow(tx pgx.Tx, logger *slog.Logger) *uow {
	q := sqlc_repository.New(tx)

	return &uow{
		tx:      tx,
		queries: q,

		bookingRepo:      repository.NewBookingRepositoryFromQueries(q, logger),
		ticketRepo:       repository.NewTicketRepositoryFromQueries(q, logger),
		bookingItemsRepo: repository.NewBookingItemsRepositoryFromQueries(q, logger),
	}
}

func (u *uow) BookingRepo() repositoryInteface.BookingRepository {
	return u.bookingRepo
}

func (u *uow) TicketRepo() repositoryInteface.TicketRepository {
	return u.ticketRepo
}

func (u *uow) BookingItemsRepo() repositoryInteface.BookingItemsRepository {
	return u.bookingItemsRepo
}

func (u *uow) Commit(ctx context.Context) error {
	return u.tx.Commit(ctx)
}

func (u *uow) Rollback(ctx context.Context) error {
	return u.tx.Rollback(ctx)
}
