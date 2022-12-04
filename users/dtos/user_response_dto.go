package dtos

import (
	"users/models"

	"github.com/google/uuid"
)

type UserResponseDto struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
}

func (user UserResponseDto) FromUser(u *models.User) UserResponseDto {
	return UserResponseDto{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}
