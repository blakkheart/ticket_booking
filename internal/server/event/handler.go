package eventapi

import (
	"net/http"
	"ticket-booking/internal/event"
	"ticket-booking/internal/httpx"
)

type Handler interface {
	CreateEvent(w http.ResponseWriter, r *http.Request) (httpx.Response, error)
	GetEvent(w http.ResponseWriter, r *http.Request) (httpx.Response, error)
	GetEvents(w http.ResponseWriter, r *http.Request) (httpx.Response, error)
	EditEvent(w http.ResponseWriter, r *http.Request) (httpx.Response, error)

	Routes(r *httpx.Router)
}

type handler struct {
	service event.Service
}

func NewHandler(service event.Service) Handler {
	return &handler{service: service}
}

func (h *handler) Routes(r *httpx.Router) {
	r.Handle(http.MethodPost, "/event/create", h.CreateEvent)
	r.Handle(http.MethodGet, "/event/{id}", h.GetEvent)
	r.Handle(http.MethodGet, "/events", h.GetEvents)
	r.Handle(http.MethodPatch, "/event/{id}", h.EditEvent)
}
