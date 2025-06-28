package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type ConnWrapper struct {
	Conn *pgx.Conn
}

func Connect() (*pgx.Conn, error) {
	url := os.Getenv("PG_URL")

	if url == "" {
		return nil, fmt.Error("PG_URL env var is not set")
	}

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, fmt.Error("Faild to connect to PG instance")
	}

	return conn, nil
}
