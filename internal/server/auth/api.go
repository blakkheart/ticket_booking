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

func (h *handler) Login(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	account, err := h.service.Login(r.Context(), "email", "1")
	fmt.Println(account)

	if err != nil {
		return httpx.NewResponse(1, http.StatusBadRequest), nil
	}

	return httpx.NewResponse(1, http.StatusOK), nil
}

func (h *handler) Authorize(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {

	//utils.GetAccountFromToken(r)
	return httpx.NewResponse(1, http.StatusOK), nil
}
