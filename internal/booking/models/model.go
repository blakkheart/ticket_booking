package models

import (
	"ticket-booking/internal/money"
	"time"
)

type BookingStatus string

const (
	Pending   BookingStatus = "pending"
	Confirmed BookingStatus = "confirmed"
	Cancelled BookingStatus = "cancelled"
	Expired   BookingStatus = "expired"
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
