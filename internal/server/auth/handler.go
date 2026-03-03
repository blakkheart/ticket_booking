package authapi

import (
	"net/http"
	"ticket-booking/internal/auth"
	"ticket-booking/internal/httpx"
)

type Handler interface {
	Login(w http.ResponseWriter, r *http.Request) error
	Authorize(w http.ResponseWriter, r *http.Request) error
	Routes(r *httpx.Router)
}

type handler struct {
	service auth.Service
}

func NewHandler(service auth.Service) Handler {
	return &handler{service: service}
}

func (h *handler) Routes(r *httpx.Router) {
	r.Handle(http.MethodPost, "/login", h.Login)
	r.Handle(http.MethodPost, "/auth", h.Authorize)
}
