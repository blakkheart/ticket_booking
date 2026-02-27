package event

import (
	"net/http"
	"ticket-booking/domain/handler/api/helper"
	"ticket-booking/domain/service"
)

type Handler interface {
	CreateEvent(w http.ResponseWriter, r *http.Request)
	GetEvent(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service service.EventService
}

func NewHandler(service service.EventService) Handler {
	return &handler{service: service}
}

func (h *handler) Routes(r *helper.Router) {
	r.Handle(http.MethodPost, "/event", h.CreateEvent)
	r.Handle(http.MethodPost, "/event/{id}", h.GetEvent)
}
