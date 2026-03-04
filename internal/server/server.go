package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"ticket-booking/internal/app"
	"ticket-booking/internal/auth"
	"ticket-booking/internal/httpx"
	"ticket-booking/internal/httpx/middleware"
	authapi "ticket-booking/internal/server/auth"
	bookingapi "ticket-booking/internal/server/booking"
	eventapi "ticket-booking/internal/server/event"
	ticketapi "ticket-booking/internal/server/ticket"
	userapi "ticket-booking/internal/server/user"

	"github.com/casbin/casbin/v2"
)

type Server interface {
	Run(ctx context.Context, addr string, authEnforcer *casbin.Enforcer, jwt *auth.JWTManager) error
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

func (s *server) Run(ctx context.Context, addr string, authEnforcer *casbin.Enforcer, jwt *auth.JWTManager) error {

	handler := middleware.Chain(
		s.router,
		middleware.RecoveryMiddleware,
		middleware.RequestIDMiddleware,
		middleware.LoggingMiddleware,
		middleware.Authorizer(authEnforcer, jwt),
	)

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	errCh := make(chan error, 1)

	go func(errCh chan<- error) {
		slog.Info("Server started", "addr", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
			return
		}
		errCh <- nil
	}(errCh)

	select {
	case <-ctx.Done():
		slog.Info("Shutdown signal received")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			return err
		}

		return nil

	case err := <-errCh:
		return err
	}

}
