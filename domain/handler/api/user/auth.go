package user

import (
	"fmt"
	"net/http"

	"ticket-booking/domain/handler/api/helper"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	account := h.service.Get(1)
	fmt.Println(account)

	token := helper.GenerateToken(account)

	status := http.StatusOK

	helper.WriteJsonResponse(w, token, status)
}

func (h *handler) Authorize(w http.ResponseWriter, r *http.Request) {

	helper.GetAccountFromToken(r)
}
