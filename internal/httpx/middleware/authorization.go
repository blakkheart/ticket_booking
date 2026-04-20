package middleware

import (
	"errors"
	"log"
	"log/slog"
	"net/http"

	"ticket-booking/internal/auth"
	"ticket-booking/internal/httpx"
	"ticket-booking/internal/user/models"

	"github.com/casbin/casbin/v2"
	"github.com/google/uuid"
)

type userInfo struct {
	Role   string
	UserID *uuid.UUID
}

func getUserInfo(r *http.Request, jwt *auth.JWTManager) (*userInfo, error) {
	role := string(models.Anonymous)
	var userID uuid.UUID

	token, tokenErr := auth.GetTokenFromPayload(r)

	if tokenErr != nil {
		slog.Info("middleware auth - failed to get token", "tokenErr", tokenErr)
	} else {

		claims, err := jwt.Parse(token)

		if err != nil {
			slog.Error(err.Error())
			return nil, errors.New("Invalid token")
		} else {
			role = claims.Role
			userID = claims.UserID
		}
	}
	return &userInfo{UserID: &userID, Role: role}, nil
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

			userInfo, err := getUserInfo(r, jwt)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := httpx.WithUserID(r.Context(), userInfo.UserID)
			r = r.WithContext(ctx)

			slog.Debug("Enforcing auth", "requestID", reqID, "role", userInfo.Role, "path", r.URL.Path, "method", r.Method)

			res, err := e.Enforce(userInfo.Role, r.URL.Path, r.Method)

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
