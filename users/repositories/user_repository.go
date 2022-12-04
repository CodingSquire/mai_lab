package repositories

import (
	model "users/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	Get(id uuid.UUID) (*model.User, error)
	GetAll() ([]model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id uuid.UUID) error
}
