package model

import "time"

type Event struct {
	ID          int
	Title       string
	Description string
	Date        time.Time
	Location    string
	TicketTypes []TicketType
}
