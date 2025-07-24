package model

type TicketType struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
	EventID     int
}
