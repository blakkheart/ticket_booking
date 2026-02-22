package middleware

import (
	"log"
	"net/http"

	"ticket-booking/domain/handler/api"
	"ticket-booking/domain/model"

	"github.com/casbin/casbin/v2"
)

func Authorizer(e *casbin.Enforcer, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var role model.Role
		account, err := api.GetAccountFromToken(r)

		if err != nil {
			role = model.Anonymous
		} else {
			role = account.Role
		}
		log.Printf("Enforcing: role=%s, path=%s, method=%s", role, r.URL.Path, r.Method)

		// Get all policies for this role
		policies, _ := e.GetFilteredPolicy(0, string(role))
		log.Printf("Policies for %s: %v", role, policies)


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
