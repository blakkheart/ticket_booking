package user

import (
	"encoding/json"
	"log"
	"net/http"

	container "ticket-booking/config/service_container"
	"ticket-booking/domain/handler/api/helper"
	accountModel "ticket-booking/models/domain/account"
)

func createAccount(w http.ResponseWriter, r *http.Request) {
	var user accountModel.AccountIn
	status := http.StatusOK

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal("Parsing Error")
	}
	//do something with user
	log.Println(user)

	u := container.ContainerService.AccountService.CreateAccount(&user)
	helper.WriteJsonResponse(w, u, status)
}
