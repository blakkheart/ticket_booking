package ticket

import (
	"ticket-booking/domain/handler/api/helper"
	"ticket-booking/domain/handler/api/ticket"
	"ticket-booking/domain/service"
	ticketrep "ticket-booking/repository/postgres/ticket"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Module interface {
	RegisterRoutes(r *helper.Router)
}

type module struct {
	Handler ticket.Handler
}

func New(db *pgxpool.Pool) Module {
	ticketRepo := ticketrep.New(db)

	ticketService := service.NewT(ticketRepo)
	handler := ticket.NewHandler(ticketService)

	return &module{
		Handler: handler,
	}
}

func (m *module) RegisterRoutes(r *helper.Router) {
	// r.Handle(http.MethodPost, "/event", m.Handler.CreateEvent)
	// r.Handle(http.MethodPost, "/event/{id}", m.Handler.GetEvent)
}
