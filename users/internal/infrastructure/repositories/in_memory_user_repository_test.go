package repositories_test

import (
	"testing"

	"users/internal/domain/models"
	"users/internal/infrastructure/repositories"

	"github.com/google/uuid"
)

func TestInMemoryUserRepository(t *testing.T) {
	repo := repositories.NewInMemoryUserRepository()

	user := &models.User{
		Username:  "username",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@mail.com",
	}

	err := repo.Create(user)
	if err != nil {
		t.Errorf("Error creating user: %s", err)
	}

	if user.ID == uuid.Nil {
		t.Errorf("User ID must be set after creating user")
	}

	_, err = repo.Get(user.ID)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	_, err = repo.Get(uuid.New())
	if err == nil {
		t.Errorf("Getting user with random ID should return an error")
	}

	err = repo.Create(user)
	if err == nil {
		t.Errorf("Expected error creating user with same id")
	}

	var users []models.User = []models.User{
		{Username: "username1", FirstName: "John", LastName: "Doe", Email: "test@test.com"},
		{Username: "username2", FirstName: "John", LastName: "Doe", Email: "test@test.com"},
	}

	for _, user := range users {
		err := repo.Create(&user)
		if err != nil {
			t.Errorf("Error creating user: %s", err)
		}
	}

	allUsers := repo.GetAll()

	if len(allUsers) != len(users)+1 {
		t.Errorf("Expected %d users, got %d", len(users)+1, len(allUsers))
	}

	err = repo.Delete(user.ID)
	if err != nil {
		t.Errorf("Error deleting user: %s", err)
	}

	_, err = repo.Get(user.ID)
	if err == nil {
		t.Errorf("Expected error getting deleted user")
	}
}
