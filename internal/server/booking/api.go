package bookingapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"ticket-booking/internal/booking/models"
	"ticket-booking/internal/httpx"
)

func (h *handler) GetBooking(
	w http.ResponseWriter,
	r *http.Request,
) (httpx.Response, error) {
	//id64, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	// userID, ok := httpx.GetUserID(r.Context())
	// if !ok {
	// 	return nil, ResolveHTTPError(errors.New("Something wrong with user"))
	// }

	// get_model := config.ContainerService.Service.GetBooking(user, event)
	get_model := 1

	return httpx.NewResponse(get_model, http.StatusOK), nil
}

func (h *handler) CreateBooking(
	w http.ResponseWriter,
	r *http.Request,
) (httpx.Response, error) {

	userID, ok := httpx.GetUserID(r.Context())
	if !ok {
		return nil, ResolveHTTPError(errors.New("Something wrong with user"))
	}

	var breq models.CreateBookingRequest
	if errorJson := json.NewDecoder(r.Body).Decode(&breq); errorJson != nil {
		return nil, httpx.ErrInvalidRequestBody
	}

	bookingCreateParams, err := models.ConvertRequestToParams(&breq)
	if err != nil {
		return nil, ResolveHTTPError(err)
	}

	newBooking, err := h.service.Create(
		r.Context(),
		bookingCreateParams,
		userID,
	)
	if err != nil {
		return nil, ResolveHTTPError(err)
	}

	response := models.CreateBookingResponse{
		BookingId:  newBooking.ID.String(),
		Status:     newBooking.Status,
		TotalPrice: newBooking.TotalPrice.String(),
	}

	return httpx.NewResponse(response, http.StatusCreated), nil
}
