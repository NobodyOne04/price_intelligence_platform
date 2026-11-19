package models

import "time"

// Keyword is a search term the crawlers track.
type Keyword struct {
	ID        int       `json:"id"`
	Value     string    `json:"value"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
