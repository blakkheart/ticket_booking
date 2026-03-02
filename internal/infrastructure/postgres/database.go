package postgres

import (
	"context"
	"fmt"
	"log"
	"ticket-booking/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConnection(dbConf *config.DBConfig) *pgxpool.Pool {
	var dsn string = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Name)

	ctx := context.Background()

	conf, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("Couldnt parse configs!")
	}

	pool, err := pgxpool.NewWithConfig(ctx, conf)

	if err != nil {
		log.Fatal(err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Failed to ping db", err)
	}

	return pool
}
