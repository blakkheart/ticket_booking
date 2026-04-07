package models

type CreateBookingRequest struct {
	EventID int64                  `json:"event_id"`
	Items   []CreateBookingItemDTO `json:"items"`
}

type CreateBookingItemDTO struct {
	TicketTypeID int64 `json:"ticket_type_id"`
	Quantity     int32 `json:"quantity"`
}

type CreateBookingResponse struct {
	BookingId  int64         `json:"booking_id"`
	Status     BookingStatus `json:"status"`
	TotalPrice string        `json:"total_price"`
}
