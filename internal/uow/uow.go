package uow

import (
	"context"
	"ticket-booking/internal/repository"
)

type UnitOfWork interface {
	BookingRepo() repository.BookingRepository
	TicketRepo() repository.TicketRepository
	BookingItemsRepo() repository.BookingItemsRepository
	PaymentIntentRepo() repository.PaymentIntentRepository
	PaymentRepo() repository.PaymentRepository

	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type UnitOfWorkManager interface {
	Begin(ctx context.Context) (UnitOfWork, error)
	Do(ctx context.Context, fn func(uow UnitOfWork) error) error
}
