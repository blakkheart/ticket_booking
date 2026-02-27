package main

import (
	"log"

	"ticket-booking/config"
	"ticket-booking/internal/modules/booking"
	"ticket-booking/internal/modules/event"
	"ticket-booking/internal/modules/ticket"
	"ticket-booking/internal/modules/user"
	"ticket-booking/repository/postgres"
	"ticket-booking/server"

	"github.com/casbin/casbin/v2"
)

func main() {

	config.ReadConfigs()

	authEnforcer, authErr := casbin.NewEnforcer("./config/casbin/auth_model.conf", "./config/casbin/policy.csv")
	if authErr != nil {
		log.Fatal(authErr)
	}

	dbPool := postgres.CreateConnection(&config.DBConfig)
	defer dbPool.Close()

	// container.InitContainerService(dbPool)

	user := user.New(dbPool)
	event := event.New(dbPool)
	ticket := ticket.New(dbPool)
	booking := booking.New(dbPool)

	s := server.NewServer(
		user,
		event,
		ticket,
		booking,
	)

	addr := "localhost:8080"
	s.Run(addr, authEnforcer)

}
