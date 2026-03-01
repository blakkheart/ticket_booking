package authapi

import (
	"fmt"
	"net/http"
	"ticket-booking/internal/httpx"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	account, err := h.service.Login("email", "1")
	fmt.Println(account)

	if err != nil {
		httpx.WriteJsonResponse(w, 1, http.StatusBadRequest)
	}

	httpx.WriteJsonResponse(w, 1, http.StatusOK)
}

func (h *handler) Authorize(w http.ResponseWriter, r *http.Request) {

	//utils.GetAccountFromToken(r)
}
