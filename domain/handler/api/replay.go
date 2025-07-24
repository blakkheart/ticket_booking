package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"ticket-booking/config"
	"ticket-booking/domain/model"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
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
	event := config.ContainerService.Service.GetEvent()

	get_model := config.ContainerService.Service.GetBooking(user, event)

	WriteJsonResponse(w, get_model)
}

func RegisterRoutes(router *http.ServeMux, prefix string) {
	router.HandleFunc(fmt.Sprintf("POST %s/try", prefix), func(w http.ResponseWriter, r *http.Request) { createUser(w, r) })
	router.HandleFunc(fmt.Sprintf("GET %s/book", prefix), func(w http.ResponseWriter, r *http.Request) { getBooking(w, r) })
}
