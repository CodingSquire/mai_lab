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

func (u *UserResponseDto) FromUser(user *models.User) UserResponseDto {
	u.ID = user.ID
	u.Username = user.Username
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	return *u
}
