package event

import (
	"ticket-booking/internal/ticket"
	"time"
)

type Event struct {
	ID int64 `json:"id"`
}

type EventType struct {
	ID          int
	Title       string
	Description string
	Date        time.Time
	Location    string
	TicketTypes []ticket.TicketType
}
