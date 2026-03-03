package ticketapi

import (
	"net/http"
	"ticket-booking/internal/httpx"
)

func (h *handler) CreateTicket(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	get_model := 1

	return httpx.NewResponse(get_model, http.StatusOK), nil
}
