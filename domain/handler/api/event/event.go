package event

import (
	"encoding/json"
	"log"
	"net/http"

	"ticket-booking/domain/handler/api/helper"
	"ticket-booking/repository"
)

func createEvent(w http.ResponseWriter, r *http.Request) {
	var event_param repository.CreateEventParams
	err := json.NewDecoder(r.Body).Decode(&event_param)
	status := http.StatusOK

	if err != nil {
		log.Fatal("Parsing Error")
	}
	log.Println(event_param)

	// event := config.ContainerService.Service.CreateEvent(event_param)
	event := 1
	helper.WriteJsonResponse(w, event, status)
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	// id64, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	status := http.StatusOK

	// if err != nil {
	// 	log.Fatal("Something wrong with service")
	// }

	// event := config.ContainerService.Service.GetEvent(id64)
	event := 1

	helper.WriteJsonResponse(w, event, status)
}
