package ticket

type Ticket struct {
	ID int64 `json:"id"`
}

type TicketType struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
	EventID     int
}
