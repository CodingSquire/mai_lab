package repository

import (
	"go_services_lab/models"
	"go_services_lab/pkg/user/repository"
)

type User interface {
	Get(id int) (models.User, error)
	Create(models.User) (int, error)
	GetAll() ([]models.User, error)
	Delete(id int) (int, error)
}

type ServiceUser struct {
	User
}

func NewServiceUser(rep *repository.UserRepository) *ServiceUser {
	return &ServiceUser{
		User: NewUserService(rep.User),
	}
}
