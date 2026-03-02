package bookingapi

import (
	"net/http"
	"ticket-booking/internal/booking"
	"ticket-booking/internal/httpx"
)

type Handler interface {
	GetBooking(w http.ResponseWriter, r *http.Request)
	Routes(r *httpx.Router)
}

type handler struct {
	service booking.Service
}

func NewHandler(service booking.Service) Handler {
	return &handler{service: service}
}

func (h *handler) Routes(r *httpx.Router) {
	// r.Handle(http.MethodPost, "/book", h.GetBooking)
}
