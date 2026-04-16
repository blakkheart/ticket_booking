package models

import (
	"ticket-booking/internal/money"
	"time"
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
	ID         int64
	AccountID  int64
	Status     BookingStatus
	ExpiresAt  *time.Time
	PaidAt     *time.Time
	TotalPrice *money.Money
}

type BookingIn struct {
	AccountID int64
	Status    BookingStatus
	ExpiresAt *time.Time
	PaidAt    *time.Time
}
