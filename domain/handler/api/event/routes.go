package event

import (
	"net/http"
	"ticket-booking/domain/handler/api/helper"
)

func Routes(r *helper.Router) {
	r.Handle(http.MethodPost, "/event", createEvent)
	r.Handle(http.MethodPost, "/event/{id}", getEvent)
}

// addHandler(router, http.MethodPost, fmt.Sprintf("%s/event", prefix), createEvent)
// addHandler(router, http.MethodGet, fmt.Sprintf("%s/event/{id}", prefix), getEvent)
