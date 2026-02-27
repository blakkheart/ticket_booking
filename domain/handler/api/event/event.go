package event

import (
	"encoding/json"
	"log"
	"net/http"

	"ticket-booking/domain/handler/api/helper"
)

func (h *handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event_param struct{}
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

func (h *handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	// id64, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	status := http.StatusOK

	// if err != nil {
	// 	log.Fatal("Something wrong with service")
	// }

	// event := config.ContainerService.Service.GetEvent(id64)
	event := 1

	helper.WriteJsonResponse(w, event, status)
}
