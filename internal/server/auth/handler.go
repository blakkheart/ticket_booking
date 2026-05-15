package authapi

import (
	"net/http"
	"ticket-booking/internal/auth"
	"ticket-booking/internal/httpx"
)

type Handler interface {
	Login(w http.ResponseWriter, r *http.Request) (httpx.Response, error)
	RefreshToken(w http.ResponseWriter, r *http.Request) (httpx.Response, error)
	Routes(r *httpx.Router)
}

type handler struct {
	service auth.Service
}

func NewHandler(service auth.Service) Handler {
	return &handler{service: service}
}

func (h *handler) Routes(r *httpx.Router) {
	r.Handle(http.MethodPost, "/auth/login", h.Login)
	r.Handle(http.MethodPost, "/auth/refresh-token", h.RefreshToken)
}
