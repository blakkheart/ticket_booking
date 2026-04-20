package models

import (
	"ticket-booking/internal/money"
	"time"

	"github.com/google/uuid"
)

type BookingStatus string

const (
	StatusPending         BookingStatus = "pending"
	StatusAwaitingPayment BookingStatus = "awaiting_payment"
	StatusPaid            BookingStatus = "paid"
	StatusCancelled       BookingStatus = "cancelled"
	StatusExpired         BookingStatus = "expired"
)

type Booking struct {
	ID         uuid.UUID
	AccountID  uuid.UUID
	Status     BookingStatus
	ExpiresAt  *time.Time
	PaidAt     *time.Time
	TotalPrice *money.Money
}

type BookingIn struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	Status    BookingStatus
	ExpiresAt *time.Time
	PaidAt    *time.Time
}

type CreateBookingParams struct {
	EventID uuid.UUID
	Items   []CreateBookingItemParams
}

type CreateBookingItemParams struct {
	TicketTypeID uuid.UUID
	Quantity     int32
}
