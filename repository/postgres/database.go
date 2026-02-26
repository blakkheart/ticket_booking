package postgres

import (
	"context"
	"fmt"
	"log"
	"ticket-booking/config"
	"ticket-booking/models/database"
	"ticket-booking/repository"

	"github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB database.DBStruct = database.DBStruct{}

func CreateConnection(dbConf *config.DBConfigStruct) *database.DBStruct {
	var dsn string = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.DBname)

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

	return &DB
}

func CreateConnectionGorm(dbConf *config.DBConfigStruct) *gorm.DB {
	var dsn string = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.DBname)

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
	fn func(*repository.Queries) error,
) error {

	tx, err := DB.Conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	q := repository.New(tx)

	if err := fn(q); err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)

}
