package auth

import (
	"fmt"
	"net/http"

	container "ticket-booking/config/service_container"
	"ticket-booking/domain/handler/api/helper"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func login(w http.ResponseWriter, r *http.Request) {
	account := container.ContainerService.AccountService.GetUser()
	fmt.Println(account)

	token := helper.GenerateToken(account)

	status := http.StatusOK

	helper.WriteJsonResponse(w, token, status)
}

func authorize(w http.ResponseWriter, r *http.Request) {

	helper.GetAccountFromToken(r)
}
