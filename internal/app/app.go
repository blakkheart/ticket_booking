package app

import (
	"ticket-booking/internal/auth"
	"ticket-booking/internal/booking"
	"ticket-booking/internal/event"
	"ticket-booking/internal/infrastructure/postgres"
	"ticket-booking/internal/ticket"
	"ticket-booking/internal/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	User    user.Service
	Event   event.Service
	Ticket  ticket.Service
	Booking booking.Service
	Auth    auth.Service
}

func NewApp(pool *pgxpool.Pool, jwt *auth.JWTManager) *App {
	userRepo := postgres.NewUserRepository(pool)
	userService := user.NewService(userRepo)

	eventRepo := postgres.NewEventRepository(pool)
	eventService := event.NewService(eventRepo)

	ticketRepo := postgres.NewTicketRepository(pool)
	ticketService := ticket.NewService(ticketRepo)

	bookingRepo := postgres.NewBookingRepository(pool)
	bookingService := booking.NewService(bookingRepo)

	authService := auth.NewService(jwt, userService)

	return &App{
		User:    userService,
		Event:   eventService,
		Ticket:  ticketService,
		Booking: bookingService,
		Auth:    authService,
	}
}
