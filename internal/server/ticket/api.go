package ticketapi

import (
	"net/http"
	"ticket-booking/internal/httpx"
)

func (h *handler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	get_model := 1

	status := http.StatusOK

	httpx.WriteJsonResponse(w, get_model, status)
}
