package repository

import (
	"go_services_lab/models"
	"go_services_lab/pkg/user/repository"
)

type UserService struct {
	rep repository.User
}

func NewUserService(rep repository.User) *UserService {
	return &UserService{rep: rep}
}

func (s *UserService) Get(id int) (models.User, error) {
	return s.rep.Get(id)
}

func (s *UserService) Create(user models.User) (int, error) {
	return s.rep.Create(user)
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.rep.GetAll()
}

func (s *UserService) Delete(id int) (int, error) {
	return s.rep.Delete(id)
}
