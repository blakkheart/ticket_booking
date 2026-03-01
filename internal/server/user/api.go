package userapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"ticket-booking/internal/httpx"
	"ticket-booking/internal/user"
)

func (h *userHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var user user.AccountIn
	status := http.StatusOK

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal("Parsing Error")
	}

	u := h.service.Create(r.Context(), &user)
	httpx.WriteJsonResponse(w, u, status)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	account := h.service.Get(1)
	fmt.Println(account)

	//token := utils.GenerateToken(account)

	status := http.StatusOK

	httpx.WriteJsonResponse(w, 1, status)
}

func (h *userHandler) Authorize(w http.ResponseWriter, r *http.Request) {

	//utils.GetAccountFromToken(r)
}
