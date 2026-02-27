package booking

import (
	"net/http"
	"ticket-booking/domain/handler/api/helper"
)

func (h *handler) GetBooking(w http.ResponseWriter, r *http.Request) {
	//id64, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	// user := config.ContainerService.Service.GetUser()
	// event := config.ContainerService.Service.GetEvent(32)

	// get_model := config.ContainerService.Service.GetBooking(user, event)
	get_model := 1

	status := http.StatusOK

	helper.WriteJsonResponse(w, get_model, status)
}
