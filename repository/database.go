package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func CreateConnectionGorm() *gorm.DB {
	var dsn string = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func WithTransaction(
	ctx context.Context,
	fn func(*Queries) error,
) error {

	tx, err := DB.Conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx)

	if err := fn(q); err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)

}
