package booking

import "errors"

var (
	BookingItemsAlreadyExists = errors.New("Booking for that ticket type exists")
)
