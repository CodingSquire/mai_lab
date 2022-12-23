package contracts

import (
	"users/internal/domain/models"

	"github.com/google/uuid"
)

// UserService is an interface for user services.
type UserService interface {
	GetUserById(id uuid.UUID) (*models.User, error)
	GetAllUsers() []models.User
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id uuid.UUID) error
}
