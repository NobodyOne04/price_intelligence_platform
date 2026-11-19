package db

import (
	"context"

	"api/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

// GetKeywords returns the enabled keywords.
func GetKeywords(ctx context.Context, pool *pgxpool.Pool) ([]models.Keyword, error) {
	rows, err := pool.Query(ctx, `
		SELECT id, value, enabled, created_at, updated_at
		FROM keywords
		WHERE enabled = TRUE
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []models.Keyword{}
	for rows.Next() {
		var k models.Keyword
		if err := rows.Scan(&k.ID, &k.Value, &k.Enabled, &k.CreatedAt, &k.UpdatedAt); err != nil {
			return nil, err
		}
		results = append(results, k)
	}
	return results, rows.Err()
}

// InsertKeyword adds an enabled keyword.
func InsertKeyword(ctx context.Context, pool *pgxpool.Pool, value string) error {
	_, err := pool.Exec(ctx, `INSERT INTO keywords (value, enabled) VALUES ($1, TRUE)`, value)
	return err
}
