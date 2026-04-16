package paymentintent

import "errors"

var (
	ErrInvalidBookingState = errors.New("Invalid booking status state")
	ErrBookingExpired      = errors.New("Booking Expired")
)
