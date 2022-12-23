// Package contracts contains all the interfaces for the application.
package contracts

import (
	model "users/internal/domain/models"

	"github.com/google/uuid"
)

// UserRepository is an interface for user repositories.
type UserRepository interface {
	Get(id uuid.UUID) (*model.User, error)
	GetAll() []model.User
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id uuid.UUID) error
}
