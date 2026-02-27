package ticket

import (
	"net/http"
	"ticket-booking/domain/handler/api/helper"
	"ticket-booking/domain/service"
)

type Handler interface {
	CreateTicket(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service service.TicketService
}

func NewHandler(service service.TicketService) Handler {
	return &handler{service: service}
}

func (h *handler) Routes(r *helper.Router) {
	r.Handle(http.MethodPost, "/ticket", h.CreateTicket)
}
