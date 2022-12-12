package user

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `json:"id ,omitempty"`
	Name         string    `json:"name"`
	Email        string    `json:"email" ,omitempty:"email"`
	Mobile       string    `json:"mobile"`
	PasswordHash string    `json:"-"`
}

type key uuid.UUID
