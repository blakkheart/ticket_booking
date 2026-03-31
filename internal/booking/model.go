package booking

import "time"

type Booking struct {
	ID        int64
	AccountID int64
	Status    string
	ExpiresAt *time.Time
	PaidAt    *time.Time
}

type BookingIn struct {
	AccountID int64
	Status    string
	ExpiresAt *time.Time
	PaidAt    *time.Time
}
