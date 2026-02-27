package user

import (
	"encoding/json"
	"log"
	"net/http"

	"ticket-booking/domain/handler/api/helper"
	accountModel "ticket-booking/models/domain/account"
)

func (h *handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var user accountModel.AccountIn
	status := http.StatusOK

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal("Parsing Error")
	}

	u := h.service.Create(r.Context(), &user)
	helper.WriteJsonResponse(w, u, status)
}
