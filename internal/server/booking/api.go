package bookingapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"ticket-booking/internal/booking"
	"ticket-booking/internal/httpx"
)

func (h *handler) GetBooking(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	//id64, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	// user := config.ContainerService.Service.GetUser()
	// event := config.ContainerService.Service.GetEvent(32)

	// get_model := config.ContainerService.Service.GetBooking(user, event)
	get_model := 1

	return httpx.NewResponse(get_model, http.StatusOK), nil
}

func (h *handler) CreateBooking(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {

	userID, ok := httpx.GetUserID(r.Context())
	if !ok {
		return nil, ResolveHTTPError(errors.New("Something wrong with user"))
	}

	var breq booking.CreateBookingRequest
	if errorJson := json.NewDecoder(r.Body).Decode(&breq); errorJson != nil {
		return nil, httpx.ErrInvalidRequestBody
	}

	newBooking, err := h.service.Create(r.Context(), &breq, userID)
	if err != nil {
		return nil, ResolveHTTPError(err)
	}

	// totalPrice, _ := money.NewMoney("1")
	response := booking.CreateBookingResponse{
		BookingId:  newBooking.ID,
		Status:     newBooking.Status,
		TotalPrice: "1",
	}

	return httpx.NewResponse(response, http.StatusCreated), nil
}
