package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DBStruct struct {
	Ctx  context.Context
	Conn *pgx.Conn
}
