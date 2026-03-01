package authapi

import (
	"ticket-booking/internal/auth"
	"ticket-booking/internal/httpx"
)

type Handler interface {
	// GetBooking(w http.ResponseWriter, r *http.Request)
	Routes(r *httpx.Router)
}

type handler struct {
	service auth.Service
}

func NewHandler(service auth.Service) Handler {
	return &handler{service: service}
}

func (h *handler) Routes(r *httpx.Router) {
}
