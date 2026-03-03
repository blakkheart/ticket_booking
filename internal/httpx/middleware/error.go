package middleware

import (
	"net/http"
	"ticket-booking/internal/httpx"
)

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)

			if errVal := r.Context().Value(httpx.ErrorKey); errVal != nil {
				err := errVal.(error)

				httpErr := httpx.ResolveHTTPError(err)
				httpx.JsonResponse(w, map[string]string{
					"code":    httpErr.Code,
					"message": httpErr.Message,
				}, httpErr.Status)
			}

		},
	)
}
