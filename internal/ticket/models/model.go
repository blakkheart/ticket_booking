package models

import "ticket-booking/internal/money"

type TicketType struct {
	ID                int64
	Name              string
	Description       *string
	Price             money.Money
	AvailableQuantity int32
	EventID           int64
}

type TicketTypeIn struct {
	Name              string
	Description       *string
	Price             money.Money
	AvailableQuantity int32
	EventID           int64
}
