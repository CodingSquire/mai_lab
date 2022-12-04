package repositories_test

import (
	"testing"

	"users/models"
	"users/repositories"

	"github.com/google/uuid"
)

func TestInMemoryUserRepository(t *testing.T) {
	creatingAndGettingUsers(t)
	deletingUsers(t)
}

func creatingAndGettingUsers(t *testing.T) {
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

	// uuid.UUID must be set after creating user
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

	allUsers, err := repo.GetAll()
	if err != nil {
		t.Errorf("Error getting all users: %s", err)
	}

	if len(allUsers) != len(users)+1 {
		t.Errorf("Expected %d users, got %d", len(users)+1, len(allUsers))
	}
}

func deletingUsers(t *testing.T) {
	//repo := repositories.NewInMemoryUserRepository()

}
