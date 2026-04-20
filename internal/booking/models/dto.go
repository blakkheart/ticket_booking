package models

type CreateBookingRequest struct {
	EventID string                 `json:"event_id"`
	Items   []CreateBookingItemDTO `json:"items"`
}

type CreateBookingItemDTO struct {
	TicketTypeID string `json:"ticket_type_id"`
	Quantity     int32  `json:"quantity"`
}

type CreateBookingResponse struct {
	BookingId  string        `json:"booking_id"`
	Status     BookingStatus `json:"status"`
	TotalPrice string        `json:"total_price"`
}
