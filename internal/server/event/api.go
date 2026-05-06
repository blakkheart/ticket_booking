package eventapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ticket-booking/internal/event/models"
	"ticket-booking/internal/httpx"

	"github.com/google/uuid"
)

func (h *handler) CreateEvent(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	var event_param *models.EventIn
	err := json.NewDecoder(r.Body).Decode(event_param)

	if err != nil {
		return nil, fmt.Errorf("Cannot decode event_params %w", err)
	}

	event, err := h.service.Create(r.Context(), event_param)

	if err != nil {
		return nil, err
	}

	return httpx.NewResponse(event, http.StatusOK), nil
}

func (h *handler) EditEvent(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	uuidString := r.PathValue("id")
	convertedUUID, err := uuid.Parse(uuidString)
	if err != nil {
		return nil, err
	}

	var event_param *models.EventUpdate
	err = json.NewDecoder(r.Body).Decode(event_param)
	if err != nil {
		return nil, fmt.Errorf("Cannot decode event_params %w", err)
	}

	event, err := h.service.UpdateById(r.Context(), convertedUUID, event_param)
	if err != nil {
		return nil, err
	}
	return httpx.NewResponse(event, http.StatusOK), nil
}

func (h *handler) GetEvent(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	uuidString := r.PathValue("id")
	convertedUUID, err := uuid.Parse(uuidString)
	if err != nil {
		return nil, err
	}

	event, err := h.service.Get(r.Context(), convertedUUID)
	if err != nil {
		return nil, err
	}
	return httpx.NewResponse(event, http.StatusOK), nil
}

func (h *handler) GetEvents(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	event := 1
	return httpx.NewResponse(event, http.StatusOK), nil
}
