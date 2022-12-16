// Package models contains the data models for the application.
package models

import "github.com/google/uuid"

// User represents a user.
type User struct {
	ID        uuid.UUID
	Username  string
	FirstName string
	LastName  string
	Email     string
	CreatedAt int64
	UpdatedAt int64
}
