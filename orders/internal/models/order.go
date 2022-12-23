package models

import "time"

type Order struct {
	ID        string
	UserID    string
	Item      string
	Address   string
	Count     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
