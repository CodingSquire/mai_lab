package repositories

import (
	"errors"
	"sync"
	"time"

	"users/models"

	"github.com/google/uuid"
)

type InMemoryUserRepository struct {
	users map[uuid.UUID]*models.User
	mu    sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[uuid.UUID]*models.User),
	}
}

func (r *InMemoryUserRepository) Get(id uuid.UUID) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *InMemoryUserRepository) GetAll() ([]models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]models.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, *user)
	}

	return users, nil
}

func (r *InMemoryUserRepository) Create(user *models.User) error {
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

func (r *InMemoryUserRepository) Update(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[user.ID]; !ok {
		return errors.New("user not found")
	}

	user.UpdatedAt = time.Now().Unix()
	r.users[user.ID] = user

	return nil
}

func (r *InMemoryUserRepository) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[id]; !ok {
		return errors.New("user does not exist")
	}

	delete(r.users, id)

	return nil
}
