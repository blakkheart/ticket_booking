package eventapi

import (
	"net/http"
	"ticket-booking/internal/event"
	"ticket-booking/internal/httpx"
)

type Handler interface {
	CreateEvent(w http.ResponseWriter, r *http.Request) error
	GetEvent(w http.ResponseWriter, r *http.Request) error
	Routes(r *httpx.Router)
}

type handler struct {
	service event.Service
}

func NewHandler(service event.Service) Handler {
	return &handler{service: service}
}

func (h *handler) Routes(r *httpx.Router) {
	r.Handle(http.MethodPost, "/event", h.CreateEvent)
	r.Handle(http.MethodPost, "/event/{id}", h.GetEvent)
}
