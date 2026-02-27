package user

import (
	"ticket-booking/domain/handler/api/helper"
	"ticket-booking/domain/handler/api/user"
	"ticket-booking/domain/service"
	accrep "ticket-booking/repository/postgres/account"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Module interface {
	RegisterRoutes(r *helper.Router)
}

type module struct {
	Handler user.Handler
}

func New(db *pgxpool.Pool) Module {
	accountRepo := accrep.New(db)

	accountService := service.New(accountRepo)
	handler := user.NewHandler(accountService)

	return &module{
		Handler: handler,
	}
}

func (m *module) RegisterRoutes(r *helper.Router) {
	m.Handler.Routes(r)
}
