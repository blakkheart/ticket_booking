package authapi

import (
	"encoding/json"
	"net/http"
	"ticket-booking/internal/httpx"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	var logReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&logReq); err != nil {
		return nil, httpx.ErrInvalidRequestBody
	}

	account, err := h.service.Login(r.Context(), logReq.Email, logReq.Password)

	if err != nil {
		return nil, err
	}

	return httpx.NewResponse(account, http.StatusOK), nil
}

