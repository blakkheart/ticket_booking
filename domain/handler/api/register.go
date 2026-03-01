package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"ticket-booking/config"
	"ticket-booking/repository"
)

func createAccount(w http.ResponseWriter, r *http.Request) {
	var user repository.Account
	status := http.StatusOK

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal("Parsing Error")
	}
	//do something with user
	log.Println(user)

	u := config.ContainerService.Service.CreateAccount(user)
	WriteJsonResponse(w, u, status)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

var err = errors.New("email or password is incorrect ")

var jwtSecretKey = []byte("super-duper-secret-key")

func login(w http.ResponseWriter, r *http.Request) {
	account := config.ContainerService.Service.GetUser()

	token := generateToken(account)

	status := http.StatusOK

	WriteJsonResponse(w, token, status)
}

func authorize(w http.ResponseWriter, r *http.Request) {

	getTokenSub(r)
}
