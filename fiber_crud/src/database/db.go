package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ConectDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "host=localhost port=5432 user=postgres password=123456789 dbname=fiber sslmode=disable")

	if err != nil {
		return nil, err
	}

	return conn, nil

}
