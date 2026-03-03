package eventapi

import (
	"encoding/json"
	"log"
	"net/http"
	"ticket-booking/internal/httpx"
)

func (h *handler) CreateEvent(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	var event_param struct{}
	err := json.NewDecoder(r.Body).Decode(&event_param)

	if err != nil {
		log.Fatal("Parsing Error")
	}
	log.Println(event_param)

	// event := config.ContainerService.Service.CreateEvent(event_param)
	event := 1
	return httpx.NewResponse(event, http.StatusOK), nil
}

func (h *handler) GetEvent(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	// id64, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	// if err != nil {
	// 	log.Fatal("Something wrong with service")
	// }

	// event := config.ContainerService.Service.GetEvent(id64)
	event := 1

	return httpx.NewResponse(event, http.StatusOK), nil
}
