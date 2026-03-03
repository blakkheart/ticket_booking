package middleware

import (
	"log"
	"net/http"

	"ticket-booking/internal/auth"
	"ticket-booking/internal/user"

	"github.com/casbin/casbin/v2"
)

func getRole(r *http.Request, jwt *auth.JWTManager) string {
	role := string(user.Anonymous)

	token, tokenErr := auth.GetTokenFromPayload(r)

	if tokenErr != nil {
		log.Printf("middleware auth - failed to get token: %v", tokenErr)
	} else {

		claims, err := jwt.Parse(token)

		if err == nil {
			role = claims.Role
		}
	}
	return role
}

func Authorizer(e *casbin.Enforcer, jwt *auth.JWTManager) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.URL.Path == "/api/register" {
				next.ServeHTTP(w, r)
				return
			}

			role := getRole(r, jwt)

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
