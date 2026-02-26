package ticket

import (
	"net/http"

	"ticket-booking/domain/handler/api/helper"
)

func createTicket(w http.ResponseWriter, r *http.Request) {
	get_model := 1

	status := http.StatusOK

	helper.WriteJsonResponse(w, get_model, status)
}
