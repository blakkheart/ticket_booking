package middleware

import (
	"log"
	"net/http"

	"ticket-booking/internal/auth"
	"ticket-booking/internal/user"

	"github.com/casbin/casbin/v2"
)

func Authorizer(e *casbin.Enforcer, jwt *auth.JWTManager) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.URL.Path == "/api/register" {
				next.ServeHTTP(w, r)
				return
			}

			var role string

			token, error := auth.GetTokenFromPayload(r)

			if error != nil {
				log.Printf("middleware auth - failed to get token: %v", error)
			} else {

				claims, err := jwt.Parse(token)

				if err != nil {
					role = string(user.Anonymous)
				} else {
					role = claims.Role
				}
			}
			log.Printf("Enforcing: role=%s, path=%s, method=%s", role, r.URL.Path, r.Method)

			res, err := e.Enforce(role, r.URL.Path, r.Method)

			if err != nil {
				log.Printf("middleware auth - enforce error: %v", err)
				http.Error(w, "Authorization error", http.StatusInternalServerError)
			}

			if !res {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
