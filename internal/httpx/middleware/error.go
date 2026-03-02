package middleware

import (
	"fmt"
	"net/http"
	"ticket-booking/internal/httpx"
)

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			fmt.Println("context")
			fmt.Println(r.Context())
			if errVal := r.Context().Value(httpx.ErrorKey); errVal != nil {
				err := errVal.(error)

				httpErr := httpx.ResolveHTTPError(err)
				httpx.WriteJsonResponse(w, map[string]string{
					"code":    httpErr.Code,
					"message": httpErr.Message,
				}, httpErr.Status)
			}

		},
	)
}
