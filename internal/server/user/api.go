package userapi

import (
	"encoding/json"
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
