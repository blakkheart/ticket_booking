package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"ticket-booking/domain/handler/api"
	"ticket-booking/repository"
)

func UnknownHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 error")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Finished in %v", time.Since(start))
	})
}

func main() {
	repository.CreateConnection()
	defer repository.DB.Conn.Close(repository.DB.Ctx)

	mux := http.NewServeMux()

	api.RegisterRoutes(mux, "/api")
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) { UnknownHandler(w, r) })

	loggedMux := loggingMiddleware(mux)

	fmt.Println("Server started")

	err := http.ListenAndServe("localhost:8080", loggedMux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}
