package bookingitems

import "ticket-booking/internal/money"

type BookingItem struct {
	ID             int64
	BookingID      int64
	TicketTypeID   int64
	Quantity       int32
	PriceAtBooking money.Money
}

type BookingItemIn struct {
	BookingID      int64
	TicketTypeID   int64
	Quantity       int32
	PriceAtBooking money.Money
}
