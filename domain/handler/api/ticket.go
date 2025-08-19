package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"ticket-booking/config"
	"ticket-booking/repository"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	var user repository.Account
	err := json.NewDecoder(r.Body).Decode(&user)

	status := http.StatusOK

	if err != nil {
		log.Fatal("Parsing Error")
		status = http.StatusBadRequest
	}
	//do something with user
	log.Println(user)

	u := config.ContainerService.Service.CreateAccount(user)
	WriteJsonResponse(w, u, status)
}

func getBooking(w http.ResponseWriter, r *http.Request) {
	//id64, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	user := config.ContainerService.Service.GetUser()
	event := config.ContainerService.Service.GetEvent(32)

	get_model := config.ContainerService.Service.GetBooking(user, event)

	status := http.StatusOK

	WriteJsonResponse(w, get_model, status)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var event_param repository.CreateEventParams
	err := json.NewDecoder(r.Body).Decode(&event_param)
	status := http.StatusOK

	if err != nil {
		log.Fatal("Parsing Error")
	}
	log.Println(event_param)

	event := config.ContainerService.Service.CreateEvent(event_param)
	WriteJsonResponse(w, event, status)
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	id64, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	status := http.StatusOK

	if err != nil {
		log.Fatal("Something wrong with service")
	}

	event := config.ContainerService.Service.GetEvent(id64)

	WriteJsonResponse(w, event, status)
}

func createTicket(w http.ResponseWriter, r *http.Request) {}
