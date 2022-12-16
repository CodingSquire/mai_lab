// Package services contains the business logic for the application.
package services

import (
	"users/internal/contracts"
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

type userService struct {
	repo contracts.UserRepository
}

// NewUserService returns a new instance of UserService.
func NewUserService(repo contracts.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// GetUserById returns a user by id. Returns an error if the user does not exist.
func (s *userService) GetUserById(id uuid.UUID) (*models.User, error) {
	user, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAllUsers returns all users.
func (s *userService) GetAllUsers() []models.User {
	return s.repo.GetAll()
}

// CreateUser creates a new user. Returns an error if the user already exists.
func (s *userService) CreateUser(user *models.User) error {
	err := s.repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser updates an existing user. Returns an error if the user does not exist.
func (s *userService) UpdateUser(user *models.User) error {
	err := s.repo.Update(user)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes an existing user. Returns an error if the user does not exist.
func (s *userService) DeleteUser(id uuid.UUID) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
