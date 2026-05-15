package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
	"ticket-booking/internal/httpx"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					log.Println(
						"panic recovered: ", rec,
						"stack_trace: ", string(debug.Stack()),
					)

					// TODO sentry

					errResp := httpx.NewResponse(
						"Something went wrong. Please try again later",
						http.StatusInternalServerError,
					)
					errResp.WriteJson(w)
				}
			}()

			next.ServeHTTP(w, r)

		},
	)
}
