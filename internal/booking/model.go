package booking

type Booking struct {
	ID int64 `json:"id"`
}

type BookingType struct {
	ID           int
	UserID       int
	EventID      int
	TicketTypeID int
	Quantity     int
}
