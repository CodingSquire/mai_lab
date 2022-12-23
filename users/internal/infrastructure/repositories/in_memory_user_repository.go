// Package repositories implements the repository pattern for the application.
package repositories

import (
	"errors"
	"sync"
	"time"

	"users/internal/contracts"
	"users/internal/domain/models"

	"github.com/google/uuid"
)

type inMemoryUserRepository struct {
	users map[uuid.UUID]*models.User
	mu    sync.RWMutex
}

// NewInMemoryUserRepository returns a new instance of UserRepository.
func NewInMemoryUserRepository() contracts.UserRepository {
	return &inMemoryUserRepository{
		users: make(map[uuid.UUID]*models.User),
	}
}

// Get returns a user by id. Returns an error if the user does not exist.
func (r *inMemoryUserRepository) Get(id uuid.UUID) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// GetAll returns all users.
func (r *inMemoryUserRepository) GetAll() []models.User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]models.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, *user)
	}

	return users
}

// Create creates a new user. Returns an error if the user already exists.
func (r *inMemoryUserRepository) Create(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[user.ID]; ok {
		return errors.New("user already exists")
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now().Unix()
	r.users[user.ID] = user

	return nil
}

// Update updates an existing user. Returns an error if the user does not exist.
func (r *inMemoryUserRepository) Update(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[user.ID]; !ok {
		return errors.New("user not found")
	}

	user.UpdatedAt = time.Now().Unix()
	r.users[user.ID] = user

	return nil
}

// Delete deletes an existing user. Returns an error if the user does not exist.
func (r *inMemoryUserRepository) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[id]; !ok {
		return errors.New("user does not exist")
	}

	delete(r.users, id)

	return nil
}
