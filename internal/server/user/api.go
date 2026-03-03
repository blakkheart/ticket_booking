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
		return httpx.JsonResponse(w, "invalid request body", http.StatusBadRequest)
	}

	u, err := h.service.Create(r.Context(), &uRequest)

	if err != nil {
		switch err {
		case user.ErrEmailAlreadyUsed:
			return httpx.JsonResponse(w, err.Error(), http.StatusConflict)

		default:
			return httpx.JsonResponse(w, "internal error", http.StatusInternalServerError)
		}
	}

	resp := user.UserResponse{
		ID:    u.ID,
		Email: u.Email,
		Role:  string(u.Role),
	}

	return httpx.JsonResponse(w, resp, status)
}
