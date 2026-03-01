package server

import (
	"fmt"
	"net/http"

	"ticket-booking/internal/app"
	"ticket-booking/internal/auth"
	"ticket-booking/internal/httpx"
	authapi "ticket-booking/internal/server/auth"
	bookingapi "ticket-booking/internal/server/booking"
	eventapi "ticket-booking/internal/server/event"
	ticketapi "ticket-booking/internal/server/ticket"
	userapi "ticket-booking/internal/server/user"
	"ticket-booking/middleware"

	"github.com/casbin/casbin/v2"
)

type Server interface {
	Run(addr string, authEnforcer *casbin.Enforcer, jwt *auth.JWTManager)
}

type server struct {
	router *http.ServeMux
}

func NewServer(a *app.App) Server {
	mux := http.NewServeMux()

	userHandler := userapi.NewHandler(a.User)
	ticketHandler := ticketapi.NewHandler(a.Ticket)
	bookingHandler := bookingapi.NewHandler(a.Booking)
	eventHandler := eventapi.NewHandler(a.Event)

	authHandler := authapi.NewHandler(a.Auth)

	httpx.RegisterRoutes(
		mux,
		"/api",
		userHandler.Routes,
		eventHandler.Routes,
		ticketHandler.Routes,
		bookingHandler.Routes,
		authHandler.Routes,
	)

	return &server{
		router: mux,
	}
}

func (s *server) Run(addr string, authEnforcer *casbin.Enforcer, jwt *auth.JWTManager) {

	handler := middleware.Chain(
		s.router,
		middleware.LoggingMiddleware,
		middleware.Authorizer(authEnforcer, jwt),
	)

	fmt.Println("Server started")

	err := http.ListenAndServe(addr, handler)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
