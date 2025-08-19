package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"ticket-booking/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func WriteJsonResponse(w http.ResponseWriter, model any, status int) {
	jData, err := json.Marshal(model)
	if err != nil {
		log.Fatal("Error in struct")
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func getToken(r *http.Request) (string, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		return "", errors.New("error in payload")
	}

	payload := strings.TrimSpace(splitToken[1])

	return payload, nil

}

func getPayload(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) { return jwtSecretKey, nil },
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		return claims, nil
	} else {
		return nil, errors.New("cant parse token")
	}
}

func getAuthorization(r *http.Request) {
	token, err := getToken(r)
	if err != nil {
		log.Fatal(err)
	}

	payload, err := getPayload(token)
	if err != nil {
		log.Fatal(err)
	}

	sub, err := payload.GetSubject()
	if err != nil {
		log.Fatal(err)
	}
	print(sub)

}

func generateToken(account repository.Account) string {
	payload := jwt.MapClaims{
		"sub": account.Email,
		"iss": "ticket-booking",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		log.Fatal("Cannot generate token")
	}

	return t
}
