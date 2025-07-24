package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "wowsp"
	dbname   = "booking"
	password = "wowsp"
)

type db struct {
	Ctx  context.Context
	Conn *pgx.Conn
}

var DB db = db{}

func CreateConnection() {
	var dsn string = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Ping(ctx); err != nil {
		log.Fatal("Failed to ping db", err)
	}

	DB.Ctx = ctx
	DB.Conn = conn

}
