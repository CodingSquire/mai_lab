package user

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `json:"id ,omitempty"`
	Name         string    `json:"name"`
	Email        string    `json:"email" ,omitempty:"email"`
	Mobile       string    `json:"mobile"`
	PasswordHash string    `json:"-"`
}

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email" ,omitempty`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type key uuid.UUID
