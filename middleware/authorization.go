package middleware

import (
	"log"
	"net/http"

	"ticket-booking/domain/handler/api"

	"github.com/casbin/casbin/v2"
)

func Authorizer(e *casbin.Enforcer, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var role string
		account, err := api.GetAccountFromToken(r)

		if err != nil {
			role = "anonymus"
		} else {
			role = "admin"
		}
		log.Print("token sub: ", account)

		res, err := e.Enforce(role, r.URL.Path, r.Method)

		if err != nil {
			log.Fatal("middleware auth: ", err)
		}
		log.Println("res: ", res)
		if !res {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
