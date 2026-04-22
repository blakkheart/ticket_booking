package repository

import (
	"context"
	"errors"
	"log/slog"
	"ticket-booking/internal/booking"
	"ticket-booking/internal/booking_items/models"
	sqlc_repository "ticket-booking/internal/db/sqlc"
	"ticket-booking/internal/money"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type bookingItemsRepository struct {
	queries *sqlc_repository.Queries
	logger  *slog.Logger
}

func NewBookingItemsRepository(
	db *pgxpool.Pool,
	logger *slog.Logger,
) *bookingItemsRepository {
	return &bookingItemsRepository{
		queries: sqlc_repository.New(db),
		logger:  logger,
	}
}

func NewBookingItemsRepositoryFromQueries(
	q *sqlc_repository.Queries,
	logger *slog.Logger,
) *bookingItemsRepository {
	return &bookingItemsRepository{
		queries: q,
		logger:  logger,
	}
}

func (repo *bookingItemsRepository) Create(
	ctx context.Context,
	b *models.BookingItemIn,
) (*models.BookingItem, error) {
	var price pgtype.Numeric
	if err := price.Scan(b.PriceAtBooking.String()); err != nil {
		return nil, err
	}

	bookingCreated, errCreate := repo.queries.CreateBookingItem(
		ctx,
		sqlc_repository.CreateBookingItemParams{
			BookingID:      b.BookingID,
			TicketTypeID:   b.TicketTypeID,
			Quantity:       b.Quantity,
			PriceAtBooking: price,
		},
	)

	if errCreate != nil {
		var pgErr *pgconn.PgError
		if errors.As(errCreate, &pgErr) {
			if pgErr.Code == "23503" && pgErr.ConstraintName == "booking_items_ticket_type_id_fkey" {
				return nil, booking.BookingItemsAlreadyExists
			}
		}
		return nil, errCreate
	}

	return repo.fromSqlcBooking(&bookingCreated), nil
}

func (repo *bookingItemsRepository) Delete(id uuid.UUID) error {
	return nil
}

func (repo *bookingItemsRepository) Get(id uuid.UUID) (*models.BookingItem, error) {
	return nil, nil
}

func (repo *bookingItemsRepository) GetMany(filter any) ([]*models.BookingItem, error) {
	return nil, nil
}

func (repo *bookingItemsRepository) fromSqlcBooking(
	b *sqlc_repository.BookingItem,
) *models.BookingItem {
	var priceStr string
	_ = b.PriceAtBooking.Scan(&priceStr)
	moneyType, _ := money.NewMoney(priceStr)

	bookingItem := &models.BookingItem{
		ID:             b.ID,
		BookingID:      b.BookingID,
		TicketTypeID:   b.TicketTypeID,
		Quantity:       b.Quantity,
		PriceAtBooking: *moneyType,
	}

	return bookingItem
}
