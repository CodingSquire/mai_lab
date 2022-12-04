package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID
	Username  string
	FirstName string
	LastName  string
	Email     string
	CreatedAt int64
	UpdatedAt int64
}
