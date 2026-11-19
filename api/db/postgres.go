package db

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect opens a connection pool from PG_URL; a pool, not one pgx.Conn, because gin
// serves requests concurrently and a single connection isn't safe for that.
func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	url := os.Getenv("PG_URL")
	if url == "" {
		return nil, errors.New("PG_URL env var is not set")
	}

	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PG instance: %w", err)
	}
	return pool, nil
}
