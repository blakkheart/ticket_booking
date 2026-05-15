package ticketapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ticket-booking/internal/httpx"
	"ticket-booking/internal/ticket/models"

	"github.com/google/uuid"
)

func (h *handler) CreateTicket(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	var ticket_param *models.TicketTypeIn
	err := json.NewDecoder(r.Body).Decode(ticket_param)

	if err != nil {
		return nil, fmt.Errorf("Cannot decode ticket_params %w", err)
	}

	ticket, err := h.service.Create(r.Context(), ticket_param)

	if err != nil {
		return nil, err
	}

	return httpx.NewResponse(ticket, http.StatusOK), nil
}

func (h *handler) GetTicket(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	uuidString := r.PathValue("id")
	convertedUUID, err := uuid.Parse(uuidString)
	if err != nil {
		return nil, err
	}
	ticket, err := h.service.Get(convertedUUID)
	if err != nil {
		return nil, err
	}
	return httpx.NewResponse(ticket, http.StatusOK), nil
}
