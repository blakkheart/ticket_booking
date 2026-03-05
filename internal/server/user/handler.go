package userapi

import (
	"net/http"
	"ticket-booking/internal/httpx"
	"ticket-booking/internal/user"
)

type UserHandler interface {
	CreateAccount(w http.ResponseWriter, r *http.Request) (httpx.Response, error)
	Routes(r *httpx.Router)
}

type userHandler struct {
	service user.Service
}

func NewHandler(service user.Service) UserHandler {
	return &userHandler{service: service}
}

func (h *userHandler) Routes(r *httpx.Router) {
	r.Handle(http.MethodPost, "/user/register", h.CreateAccount)

}
