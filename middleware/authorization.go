package middleware

import (
	"log"
	"net/http"

	"ticket-booking/domain/handler/api/helper"
	account_model "ticket-booking/models/domain/account"

	"github.com/casbin/casbin/v2"
)

func Authorizer(e *casbin.Enforcer) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var role account_model.Role
			account, err := helper.GetAccountFromToken(r)

			if err != nil {
				role = account_model.Anonymous
			} else {
				role = account.Role
			}
			log.Printf("Enforcing: role=%s, path=%s, method=%s", role, r.URL.Path, r.Method)

			res, err := e.Enforce(string(role), r.URL.Path, r.Method)

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
}
