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

func (h *handler) Login(w http.ResponseWriter, r *http.Request) error {
	account, err := h.service.Login(r.Context(), "email", "1")
	fmt.Println(account)

	if err != nil {
		return httpx.JsonResponse(w, 1, http.StatusBadRequest)
	}

	return httpx.JsonResponse(w, 1, http.StatusOK)
}

func (h *handler) Authorize(w http.ResponseWriter, r *http.Request) error {

	//utils.GetAccountFromToken(r)
	return nil
}
