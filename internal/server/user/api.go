package userapi

import (
	"encoding/json"
	"net/http"
	"ticket-booking/internal/httpx"
	"ticket-booking/internal/user"
)

func (h *userHandler) CreateAccount(w http.ResponseWriter, r *http.Request) error {
	var uRequest user.CreateUserRequest
	status := http.StatusOK

	if err := json.NewDecoder(r.Body).Decode(&uRequest); err != nil {
		httpx.WriteJsonResponse(w, "invalid request body", http.StatusBadRequest)
		return nil
	}

	u, err := h.service.Create(r.Context(), &uRequest)

	if err != nil {
		switch err {
		case user.ErrEmailAlreadyUsed:
			httpx.WriteJsonResponse(w, err.Error(), http.StatusConflict)

		default:
			httpx.WriteJsonResponse(w, "internal error", http.StatusInternalServerError)
			return nil
		}
	}

	resp := user.UserResponse{
		ID:    u.ID,
		Email: u.Email,
		Role:  string(u.Role),
	}

	httpx.WriteJsonResponse(w, resp, status)
	return nil
}
