package ticketapi

import (
	"net/http"
	"ticket-booking/internal/httpx"
)

func (h *handler) CreateTicket(w http.ResponseWriter, r *http.Request) error {
	get_model := 1

	status := http.StatusOK

	return httpx.JsonResponse(w, get_model, status)
}
