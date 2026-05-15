package bookingapi

import (
	"errors"
	"net/http"
	"ticket-booking/internal/booking"
	"ticket-booking/internal/httpx"
)

var errorMap = map[error]httpx.HTTPError{
	booking.BookingItemsAlreadyExists: {
		Status:  http.StatusConflict,
		Message: "booking already exists",
		Code:    "BOOKING_ALREADY_EXISTS",
	},
}

func ResolveHTTPError(err error) error {
	for target, mapped := range errorMap {
		if errors.Is(err, target) {
			return &mapped
		}
	}

	return err
}
