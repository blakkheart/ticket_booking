package app

import (
	"log/slog"
	"ticket-booking/internal/auth"
	"ticket-booking/internal/booking"
	bookingitems "ticket-booking/internal/booking_items"
	"ticket-booking/internal/event"
	"ticket-booking/internal/infrastructure/postgres/repository"
	"ticket-booking/internal/ticket"
	"ticket-booking/internal/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	User         user.Service
	Event        event.Service
	Ticket       ticket.Service
	BookingItems bookingitems.Service
	Booking      booking.Service
	Auth         auth.Service
}

func NewApp(pool *pgxpool.Pool, jwt *auth.JWTManager, logger *slog.Logger) *App {
	userRepo := repository.NewUserRepository(pool, logger)
	userService := user.NewService(userRepo, logger)

	eventRepo := repository.NewEventRepository(pool, logger)
	eventService := event.NewService(eventRepo, logger)

	ticketRepo := repository.NewTicketRepository(pool, logger)
	ticketService := ticket.NewService(ticketRepo, logger)

	bookingItemsRepo := repository.NewBookingItemsRepository(pool, logger)
	bookingItemsService := bookingitems.NewService(bookingItemsRepo, logger)

	bookingRepo := repository.NewBookingRepository(pool, logger)
	bookingService := booking.NewService(bookingRepo, eventRepo, bookingItemsRepo, ticketRepo, logger)

	authRepo := repository.NewAuthRepository(pool, logger)
	authService := auth.NewService(authRepo, jwt, userService, logger)

	return &App{
		User:         userService,
		Event:        eventService,
		Ticket:       ticketService,
		BookingItems: bookingItemsService,
		Booking:      bookingService,
		Auth:         authService,
	}
}
