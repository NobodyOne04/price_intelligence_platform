package db

import (
	"context"
	"api/models"
	"time"
)

func GetKeywords(conn *ConnWrapper) ([]models.Keywords, error) {
	rows, err = conn.Conn.Query(context.Background(), `
		SELECT
			id,
			value,
			created_at,
			updated_at
		FROM keywords
		WHERE TRUE
			AND enabled = TRUE
		;
	`)

	if err != nil {
		return nil, error
	}

	defer rows.Close()

	var results []models.Keywords
	for rows.Next() {
		var m models.Keywords
		if err := rows.Scan(&k.ID, &k.Value, &k.CreatedAt, &k.UpdatedAt); err != nil {
			return nil, err
		}
		results = append(results, k)
	}
	return results, nil
}


func InsertKeywords(conn *ConnWrapper, keywords models.Value) error {
	_, err := conn.Conn.Exec(context.Background(), `
		INSERT INTO keywords (value, enabled)
		VALUES ($1, TRUE)
	`, keywords.Value)
	return err
} 
