package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"ticket-booking/config"
	"ticket-booking/repository"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	var user repository.Account
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal("Parsing Error")
	}
	//do something with user
	log.Println(user)

	u := config.ContainerService.Service.CreateAccount(user)
	WriteJsonResponse(w, u)
}

func getBooking(w http.ResponseWriter, r *http.Request) {
	user := config.ContainerService.Service.GetUser()
	event := config.ContainerService.Service.GetEvent(32)

	get_model := config.ContainerService.Service.GetBooking(user, event)

	WriteJsonResponse(w, get_model)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var event_param repository.CreateEventParams
	err := json.NewDecoder(r.Body).Decode(&event_param)
	if err != nil {
		log.Fatal("Parsing Error")
	}
	log.Println(event_param)

	event := config.ContainerService.Service.CreateEvent(event_param)
	WriteJsonResponse(w, event)
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	id64, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	if err != nil {
		log.Fatal("Something wrong with service")
	}

	event := config.ContainerService.Service.GetEvent(id64)

	WriteJsonResponse(w, event)
}

func createTicket(w http.ResponseWriter, r *http.Request) {}

func RegisterRoutes(router *http.ServeMux, prefix string) {
	router.HandleFunc(fmt.Sprintf("POST %s/try", prefix), func(w http.ResponseWriter, r *http.Request) { createUser(w, r) })
	router.HandleFunc(fmt.Sprintf("GET %s/book", prefix), func(w http.ResponseWriter, r *http.Request) { getBooking(w, r) })
	router.HandleFunc(fmt.Sprintf("POST %s/event", prefix), func(w http.ResponseWriter, r *http.Request) { createEvent(w, r) })
	router.HandleFunc(fmt.Sprintf("GET %s/event/{id}", prefix), func(w http.ResponseWriter, r *http.Request) { getEvent(w, r) })
	router.HandleFunc(fmt.Sprintf("POST %s/ticket", prefix), func(w http.ResponseWriter, r *http.Request) { createTicket(w, r) })

}
