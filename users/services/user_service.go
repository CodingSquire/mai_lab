package services

import (
	"users/models"
	"users/repositories"

	"github.com/google/uuid"
)

type UserService interface {
	GetUserById(id uuid.UUID) (*models.User, error)
	GetAllUsers() []models.User
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetUserById(id uuid.UUID) (*models.User, error) {
	user, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetAllUsers() []models.User {
	return s.repo.GetAll()
}

func (s *userService) CreateUser(user *models.User) error {
	err := s.repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) UpdateUser(user *models.User) error {
	err := s.repo.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
