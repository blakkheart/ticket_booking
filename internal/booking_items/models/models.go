package models

import (
	"ticket-booking/internal/money"

	"github.com/google/uuid"
)

type BookingItem struct {
	ID             uuid.UUID
	BookingID      uuid.UUID
	TicketTypeID   uuid.UUID
	Quantity       int32
	PriceAtBooking money.Money
}

type BookingItemIn struct {
	BookingID      uuid.UUID
	TicketTypeID   uuid.UUID
	Quantity       int32
	PriceAtBooking money.Money
}
