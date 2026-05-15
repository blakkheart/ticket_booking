package authapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"ticket-booking/internal/httpx"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	var logReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&logReq); err != nil {
		return nil, httpx.ErrInvalidRequestBody
	}

	tokens, err := h.service.Login(r.Context(), logReq.Email, logReq.Password)

	if err != nil {
		return nil, err
	}

	return httpx.NewResponse(tokens, http.StatusOK), nil
}

func (h *handler) RefreshToken(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	ctx := r.Context()

	var token RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		return nil, httpx.ErrInvalidRequestBody
	}

	claims, err := h.service.ParseToken(token.RefreshToken)
	if err != nil {
		return nil, err
	}
	existedToken, err := h.service.GetTokenByUserID(ctx, claims.UserID)

	if err != nil {
		return nil, err
	}
	if existedToken.Revoked {
		return nil, errors.New("Token revoked")
	}
	if existedToken.TokenHash != token.RefreshToken {
		return nil, errors.New("Wrong token")
	}
	if time.Now().After(existedToken.ExpiresAt) {
		return nil, errors.New("Token expired")
	}

	newToken, err := h.service.GenerateTokenPair(claims.UserID, claims.Role)
	if err != nil {
		return nil, err
	}

	return httpx.NewResponse(newToken, http.StatusOK), nil
}
