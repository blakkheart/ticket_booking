package userapi

import (
	"encoding/json"
	"net/http"
	"ticket-booking/internal/httpx"
	"ticket-booking/internal/user"
)

func (h *userHandler) CreateAccount(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	var uRequest user.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&uRequest); err != nil {
		return nil, httpx.ErrInvalidRequestBody
	}

	u, err := h.service.Create(r.Context(), &uRequest)

	if err != nil {
		return nil, ResolveHTTPError(err)
	}

	u_resp := user.UserResponse{
		ID:    u.ID,
		Email: u.Email,
		Role:  string(u.Role),
	}

	return httpx.NewResponse(u_resp, http.StatusOK), nil
}
