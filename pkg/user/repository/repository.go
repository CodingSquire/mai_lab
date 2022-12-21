package repository

import (
	"go_services_lab/models"

	"github.com/jmoiron/sqlx"
)

type User interface {
	Get(id int) (models.User, error)
	Create(models.User) (int, error)
	GetAll() ([]models.User, error)
	Delete(id int) (int, error)
}

type UserRepository struct {
	User
}

func NewRepositoryUser(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		User: NewUserPostgres(db),
	}
}
