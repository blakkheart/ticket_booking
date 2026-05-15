package middleware

import (
	"log"
	"net/http"
	"ticket-booking/internal/httpx"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req_id, ok := httpx.GetRequestID(r.Context())
		if !ok {
			req_id = "UNKNOWN"
		}

		start := time.Now()
		log.Printf("[req_id=%s] started %s %s", req_id, r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("[req_id=%s] finished in %v", req_id, time.Since(start))
	})
}
