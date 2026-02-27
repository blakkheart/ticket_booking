package user

import (
	"net/http"
	"ticket-booking/domain/handler/api/helper"
	"ticket-booking/domain/service"
)

type Handler interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Authorize(w http.ResponseWriter, r *http.Request)
	Routes(r *helper.Router)
}

type handler struct {
	service service.UserService
}

func NewHandler(service service.UserService) Handler {
	return &handler{service: service}
}

func (h *handler) Routes(r *helper.Router) {
	r.Handle(http.MethodPost, "/register", h.CreateAccount)
	r.Handle(http.MethodPost, "/login", h.Login)
	r.Handle(http.MethodPost, "/auth", h.Authorize)
}
