package middleware

import (
	"errors"
	"log"
	"log/slog"
	"net/http"

	"ticket-booking/internal/auth"
	"ticket-booking/internal/httpx"
	"ticket-booking/internal/user"

	"github.com/casbin/casbin/v2"
)

func getRole(r *http.Request, jwt *auth.JWTManager) (string, error) {
	role := string(user.Anonymous)

	token, tokenErr := auth.GetTokenFromPayload(r)

	if tokenErr != nil {
		slog.Info("middleware auth - failed to get token", "tokenErr", tokenErr)
	} else {

		claims, err := jwt.Parse(token)

		if err != nil {
			slog.Error(err.Error())
			return "", errors.New("Invalid token")
		} else {
			role = claims.Role
		}
	}
	return role, nil
}

func Authorizer(e *casbin.Enforcer, jwt *auth.JWTManager) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.URL.Path == "/api/register" {
				next.ServeHTTP(w, r)
				return
			}
			reqID, ok := httpx.GetRequestID(r.Context())
			if !ok {
				reqID = "UNKNOWN"
			}

			role, err := getRole(r, jwt)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			slog.Debug("Enforcing auth", "requestID", reqID, "role", role, "path", r.URL.Path, "method", r.Method)

			res, err := e.Enforce(role, r.URL.Path, r.Method)

			if err != nil {
				log.Printf("middleware auth - enforce error: %v", err)
				http.Error(w, "Authorization error", http.StatusInternalServerError)
			}
			slog.Debug("Enforced resolution", "resolution", res)

			if !res {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
