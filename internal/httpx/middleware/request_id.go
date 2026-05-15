package middleware

import (
	"net/http"
	"ticket-booking/internal/httpx"

	"github.com/google/uuid"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := uuid.NewString()

		ctx := httpx.WithRequstID(r.Context(), id)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
