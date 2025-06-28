package models

import "time"

type Keyword struct {
	ID        int
	Value     string
	Enabled   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
