package event

import (
	"net/http"
	"ticket-booking/domain/handler/api/event"
	"ticket-booking/domain/handler/api/helper"
	"ticket-booking/domain/service"
	eventrep "ticket-booking/repository/postgres/event"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Module interface {
	RegisterRoutes(r *helper.Router)
}

type module struct {
	Handler event.Handler
}

func New(db *pgxpool.Pool) Module {
	eventRepo := eventrep.New(db)

	eventService := service.NewE(eventRepo)
	handler := event.NewHandler(eventService)

	return &module{
		Handler: handler,
	}
}

func (m *module) RegisterRoutes(r *helper.Router) {
	r.Handle(http.MethodPost, "/event", m.Handler.CreateEvent)
	r.Handle(http.MethodPost, "/event/{id}", m.Handler.GetEvent)
}
