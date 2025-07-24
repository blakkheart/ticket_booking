package model

type Booking struct {
	ID           int
	UserID       int
	EventID      int
	TicketTypeID int
	Quantity     int
}
