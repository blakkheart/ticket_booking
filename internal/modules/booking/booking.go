package booking

import (
	"ticket-booking/domain/handler/api/booking"
	"ticket-booking/domain/handler/api/helper"
	"ticket-booking/domain/service"
	bookrep "ticket-booking/repository/postgres/booking"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Module interface {
	RegisterRoutes(r *helper.Router)
}

type module struct {
	Handler booking.Handler
}

func New(db *pgxpool.Pool) Module {
	bookingRepo := bookrep.New(db)

	bookingService := service.NewB(bookingRepo)
	handler := booking.NewHandler(bookingService)

	return &module{
		Handler: handler,
	}
}

func (m *module) RegisterRoutes(r *helper.Router) {
	// r.Handle(http.MethodPost, "/event", m.Handler.CreateEvent)
	// r.Handle(http.MethodPost, "/event/{id}", m.Handler.GetEvent)
}
