package api

import (
	"net/http"

	"ticket-booking/domain/handler/api/auth"
	"ticket-booking/domain/handler/api/booking"
	"ticket-booking/domain/handler/api/event"
	"ticket-booking/domain/handler/api/ticket"
	"ticket-booking/domain/handler/api/user"

	"ticket-booking/domain/handler/api/helper"
)

func RegisterRoutes(mux *http.ServeMux, prefix string) {
	router := helper.NewRouter(mux, prefix)

	router.Include(event.Routes)
	router.Include(auth.Routes)
	router.Include(ticket.Routes)
	router.Include(user.Routes)
	router.Include(booking.Routes)

}
