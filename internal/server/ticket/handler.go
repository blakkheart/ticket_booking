package ticketapi

import (
	"net/http"
	"ticket-booking/internal/httpx"
	"ticket-booking/internal/ticket"
)

type Handler interface {
	CreateTicket(w http.ResponseWriter, r *http.Request) error
	Routes(r *httpx.Router)
}

type handler struct {
	service ticket.Service
}

func NewHandler(service ticket.Service) Handler {
	return &handler{service: service}
}

func (h *handler) Routes(r *httpx.Router) {
	r.Handle(http.MethodPost, "/ticket", h.CreateTicket)
}
