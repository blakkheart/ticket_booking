package model_interface

import (
	"ticket-booking/models/domain/booking"
)

type IBookingRepository interface {
	IBaseRepository[booking.Booking, booking.Booking]
}
