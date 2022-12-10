// Package dtos contains the data transfer objects used in the application.
package dtos

import (
	"users/models"
)

// UserRequestDto represents a user request.
type UserRequestDto struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// ToUser converts the user request dto to a user model.
func (user UserRequestDto) ToUser() *models.User {
	return &models.User{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
