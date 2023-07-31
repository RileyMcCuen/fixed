package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context) (*pgx.Conn, error) {
	dbUrl := "postgres://postgres:G5fK$7I86$BR@test.cb0ptwfks8sw.us-east-1.rds.amazonaws.com:5432/test"
	return pgx.Connect(ctx, dbUrl)
}
