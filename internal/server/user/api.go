package userapi

import (
	"encoding/json"
	"net/http"
	"ticket-booking/internal/httpx"
	"ticket-booking/internal/user/models"

	"github.com/google/uuid"
)

func (h *userHandler) CreateAccount(w http.ResponseWriter, r *http.Request) (httpx.Response, error) {
	var uRequest models.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&uRequest); err != nil {
		return nil, httpx.ErrInvalidRequestBody
	}

	userModel := &models.CreateUserParams{
		ID:       uuid.New(),
		Name:     uRequest.Name,
		Email:    uRequest.Email,
		Password: uRequest.Password,
	}

	u, err := h.service.Create(r.Context(), userModel)

	if err != nil {
		return nil, ResolveHTTPError(err)
	}

	u_resp := models.UserResponse{
		ID:    u.ID.String(),
		Email: u.Email,
		Role:  string(u.Role),
	}

	return httpx.NewResponse(u_resp, http.StatusOK), nil
}
