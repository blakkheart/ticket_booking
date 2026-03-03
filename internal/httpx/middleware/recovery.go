package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"ticket-booking/internal/httpx"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					fmt.Println(
						"panic recovered: ", rec,
						"stack_trace: ", string(debug.Stack()),
					)

					// TODO sentry

					httpx.WriteJsonResponse(w, "internal error", http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)

		},
	)
}
