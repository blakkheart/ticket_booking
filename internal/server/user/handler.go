package userapi

import (
	"net/http"
	"ticket-booking/internal/httpx"
	"ticket-booking/internal/user"
)

type UserHandler interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Authorize(w http.ResponseWriter, r *http.Request)
	Routes(r *httpx.Router)
}

type userHandler struct {
	service user.Service
}

func NewHandler(service user.Service) UserHandler {
	return &userHandler{service: service}
}

func (h *userHandler) Routes(r *httpx.Router) {
	r.Handle(http.MethodPost, "/register", h.CreateAccount)
	r.Handle(http.MethodPost, "/login", h.Login)
	r.Handle(http.MethodPost, "/auth", h.Authorize)
}
