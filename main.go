package main

import (
	"log"
	"time"

	"ticket-booking/config"
	"ticket-booking/internal/app"
	"ticket-booking/internal/auth"
	"ticket-booking/internal/infrastructure/postgres"
	"ticket-booking/internal/server"

	"github.com/casbin/casbin/v2"
)

func main() {

	config.ReadConfigs()

	authEnforcer, authErr := casbin.NewEnforcer(
		"./config/casbin/auth_model.conf",
		"./config/casbin/policy.csv",
	)
	if authErr != nil {
		log.Fatal(authErr)
	}

	jwt := auth.NewJWTManager(config.AuthConfig.SecretKey, "test", time.Duration(1000000000000))

	dbPool := postgres.CreateConnection(&config.DBConfig)
	defer dbPool.Close()

	app := app.NewApp(dbPool, jwt)
	s := server.NewServer(app)

	s.Run(config.ServerConfig.SiteHost, authEnforcer, jwt)

}
