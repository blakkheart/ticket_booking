package main

import (
	"context"
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

	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	config.InitConfigs()

	authEnforcer, authErr := casbin.NewEnforcer(
		"./internal/config/casbin/auth_model.conf",
		"./internal/config/casbin/policy.csv",
	)
	if authErr != nil {
		slog.Error("Error occured while initializing authentication", "error", authErr)
	}

	jwt := auth.NewJWTManager(config.AppConfigs.Auth.SecretKey, "test", 24*time.Hour, 24*90*time.Hour)

	dbPool := postgres.CreateConnection(&config.AppConfigs.DB)
	defer dbPool.Close()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	app := app.NewApp(dbPool, jwt, logger)
	s := server.NewServer(app, logger)

	if err := s.Run(
		ctx,
		config.AppConfigs.Server.GetAddress(),
		authEnforcer,
		jwt,
	); err != nil {
		logger.Error("Server failed", "error", err)
	}

	slog.Info("Application stopped")
}
