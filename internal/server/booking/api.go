package bookingapi

import (
	"net/http"
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
