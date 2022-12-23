package storage

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
	"mai_lab/internal/domain/models"
)

// MemoryRepository fulfills the Storage interface
type memoryStorage struct {
	users map[uuid.UUID]models.User
	sync.Mutex
}

func NewMemoryStorage() Storage {
	return &memoryStorage{
		users: make(map[uuid.UUID]models.User),
	}
}

// the function of verifying the existence of a user with this key
func (s *memoryStorage) checkUserExist(key uuid.UUID) bool {
	if _, ok := s.users[key]; ok {
		return true
	}
	return false
}

// Create add will add a new user to the repository
func (s *memoryStorage) Create(ctx context.Context, user models.User) error {
	s.Lock()
	defer s.Unlock()
	if !s.checkUserExist(user.ID) {
		s.users[user.ID] = user
		log.Println("create user ", user.Name)
		return nil
	}
	return fmt.Errorf("failed to create user")
}

// GetUser finds a user by ID
func (s *memoryStorage) GetUser(ctx context.Context, id uuid.UUID) (models.User, error) {

	if exist := s.checkUserExist(id); exist {
		return s.users[id], nil
	} else {
		return models.User{}, fmt.Errorf("failed to find user by id: %s", id)
	}
}

// GetAll return users slice
func (s *memoryStorage) GetAll(ctx context.Context) ([]models.User, error) {
	all := make([]models.User, 0, len(s.users))
	for _, value := range s.users {
		all = append(all, value)
	}
	return all, nil
}

// Update will replace an existing user information with the new user information
func (s *memoryStorage) Update(ctx context.Context, user models.User) error {
	s.Lock()
	defer s.Unlock()
	if s.checkUserExist(user.ID) {
		s.users[user.ID] = user
		return nil
	}
	return fmt.Errorf("failed to update, user does not exist")
}

func (s *memoryStorage) Delete(ctx context.Context, id uuid.UUID) error {
	s.Lock()
	defer s.Unlock()
	if s.checkUserExist(id) {
		delete(s.users, id)
		return nil
	}
	return fmt.Errorf("failed to delete, user does not exist")
}
