package models

import (
	"ticket-booking/internal/money"

	"github.com/google/uuid"
)

type TicketType struct {
	ID                uuid.UUID
	Name              string
	Description       *string
	Price             money.Money
	AvailableQuantity int32
	EventID           uuid.UUID
}

type TicketTypeIn struct {
	ID                uuid.UUID
	Name              string
	Description       *string
	Price             money.Money
	AvailableQuantity int32
	EventID           uuid.UUID
}
