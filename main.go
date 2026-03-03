package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ticket-booking/internal/app"
	"ticket-booking/internal/auth"
	"ticket-booking/internal/config"
	"ticket-booking/internal/infrastructure/postgres"
	"ticket-booking/internal/server"

	"github.com/casbin/casbin/v2"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	config.InitConfigs()

	authEnforcer, authErr := casbin.NewEnforcer(
		"./internal/config/casbin/auth_model.conf",
		"./internal/config/casbin/policy.csv",
	)
	if authErr != nil {
		log.Fatal(authErr)
	}

	jwt := auth.NewJWTManager(config.AppConfigs.Auth.SecretKey, "test", 24*time.Hour)

	dbPool := postgres.CreateConnection(&config.AppConfigs.DB)
	defer dbPool.Close()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	app := app.NewApp(dbPool, jwt)
	s := server.NewServer(app)

	if err := s.Run(
		ctx,
		config.AppConfigs.Server.GetAddress(),
		authEnforcer,
		jwt,
	); err != nil {
		log.Fatal("Server failed: %v", err)
	}

	slog.Info("Application stopped")
}
