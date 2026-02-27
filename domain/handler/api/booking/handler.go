package booking

import (
	"net/http"
	"ticket-booking/domain/handler/api/helper"
	"ticket-booking/domain/service"
)

type Handler interface {
	GetBooking(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service service.BookingService
}

func NewHandler(service service.BookingService) Handler {
	return &handler{service: service}
}

func (h *handler) Routes(r *helper.Router) {
	r.Handle(http.MethodPost, "/book", h.GetBooking)
}
