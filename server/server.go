package server

import (
	"fmt"
	"net/http"
	"ticket-booking/internal/modules/booking"
	"ticket-booking/internal/modules/event"
	"ticket-booking/internal/modules/ticket"
	"ticket-booking/internal/modules/user"
	"ticket-booking/middleware"

	"github.com/casbin/casbin/v2"
)

type Server interface {
	Run(addr string, authEnforcer *casbin.Enforcer)
}

type server struct {
	router *http.ServeMux
}

func NewServer(
	userModule user.Module,
	eventModule event.Module,
	ticketModule ticket.Module,
	bookingModule booking.Module,
) Server {
	mux := http.NewServeMux()

	RegisterRoutes(
		mux,
		"/api",
		userModule.RegisterRoutes,
		eventModule.RegisterRoutes,
		ticketModule.RegisterRoutes,
		bookingModule.RegisterRoutes,
	)

	return &server{
		router: mux,
	}
}

func (s *server) Run(addr string, authEnforcer *casbin.Enforcer) {

	handler := middleware.Chain(
		s.router,
		middleware.LoggingMiddleware,
		middleware.Authorizer(authEnforcer),
	)

	fmt.Println("Server started")

	err := http.ListenAndServe(addr, handler)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
