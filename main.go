package main

import (
	"fmt"
	"log"
	"net/http"

	"ticket-booking/config"
	container "ticket-booking/config/service_container"
	"ticket-booking/domain/handler/api"
	"ticket-booking/middleware"
	"ticket-booking/repository/postgres"

	"github.com/casbin/casbin/v2"
)

func main() {

	config.ReadConfigs()

	authEnforcer, authErr := casbin.NewEnforcer("./config/casbin/auth_model.conf", "./config/casbin/policy.csv")
	if authErr != nil {
		log.Fatal(authErr)
	}

	dbConn := postgres.CreateConnection(&config.DBConfig)
	defer postgres.DB.Conn.Close(postgres.DB.Ctx)

	container.InitContainerService(dbConn)

	mux := http.NewServeMux()

	api.RegisterRoutes(mux, "/api")

	fmt.Println("Server started")

	err := http.ListenAndServe("localhost:8080", middleware.LoggingMiddleware(middleware.Authorizer(authEnforcer, mux)))
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}
