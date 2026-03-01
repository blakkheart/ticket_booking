package api

import (
	"fmt"
	"net/http"
)

func RegisterRoutes(router *http.ServeMux, prefix string) {
	addHandler(router, http.MethodPost, fmt.Sprintf("%s/register", prefix), createAccount)
	addHandler(router, http.MethodGet, fmt.Sprintf("%s/book", prefix), getBooking)
	addHandler(router, http.MethodPost, fmt.Sprintf("%s/event", prefix), createEvent)
	addHandler(router, http.MethodGet, fmt.Sprintf("%s/event/{id}", prefix), getEvent)
	addHandler(router, http.MethodPost, fmt.Sprintf("%s/ticket", prefix), createTicket)
	addHandler(router, http.MethodPost, fmt.Sprintf("%s/login", prefix), login)
	addHandler(router, http.MethodGet, fmt.Sprintf("%s/auth", prefix), authorize)
}

type handlerFunc func(w http.ResponseWriter, r *http.Request)

func addHandler(router *http.ServeMux, method string, path string, handlerFunc handlerFunc) {
	router.HandleFunc(fmt.Sprintf("%s %s", method, path), func(w http.ResponseWriter, r *http.Request) { handlerFunc(w, r) })
}
