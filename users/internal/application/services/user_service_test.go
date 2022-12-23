package services_test

import (
	"errors"
	"testing"
	"users/internal/application/services"
	"users/internal/contracts/mocks"
	"users/internal/domain/models"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestUserService(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()
	testGetUserById_UserExist(t, mock)
	testGetUserById_UserNotExist(t, mock)
	testGetAllUsers(t, mock)
	testCreateUser(t, mock)
	testCreateUser_UserAlreadyExist(t, mock)
	testUpdateUser(t, mock)
	testUpdateUser_UserNotExist(t, mock)
	testDeleteUser(t, mock)
	testDeleteUser_UserNotExist(t, mock)
}

func testDeleteUser_UserNotExist(t *testing.T, mock *gomock.Controller) {
	var notExistingUserID = uuid.New()
	mockUserRepository := mocks.NewMockUserRepository(mock)
	mockUserRepository.EXPECT().Delete(notExistingUserID).Return(errors.New("user not found"))
	userService := services.NewUserService(mockUserRepository)
	err := userService.DeleteUser(notExistingUserID)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func testDeleteUser(t *testing.T, mock *gomock.Controller) {
	mockUserRepository := mocks.NewMockUserRepository(mock)
	mockUserRepository.EXPECT().Delete(gomock.Any()).Return(nil)
	userService := services.NewUserService(mockUserRepository)
	err := userService.DeleteUser(uuid.New())
	if err != nil {
		t.Errorf("Error deleting user: %s", err)
	}
}

func testUpdateUser_UserNotExist(t *testing.T, mock *gomock.Controller) {
	mockUserRepository := mocks.NewMockUserRepository(mock)
	mockUserRepository.EXPECT().Update(gomock.Any()).Return(errors.New("user not found"))
	userService := services.NewUserService(mockUserRepository)
	err := userService.UpdateUser(&models.User{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func testUpdateUser(t *testing.T, mock *gomock.Controller) {
	mockUserRepository := mocks.NewMockUserRepository(mock)
	mockUserRepository.EXPECT().Update(gomock.Any()).Return(nil)
	userService := services.NewUserService(mockUserRepository)
	err := userService.UpdateUser(&models.User{})
	if err != nil {
		t.Errorf("Error updating user: %s", err)
	}
}

func testCreateUser_UserAlreadyExist(t *testing.T, mock *gomock.Controller) {
	mockUserRepository := mocks.NewMockUserRepository(mock)
	mockUserRepository.EXPECT().Create(gomock.Any()).Return(errors.New("user already exist"))
	userService := services.NewUserService(mockUserRepository)
	err := userService.CreateUser(&models.User{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func testCreateUser(t *testing.T, mock *gomock.Controller) {
	mockUserRepository := mocks.NewMockUserRepository(mock)
	mockUserRepository.EXPECT().Create(gomock.Any()).Return(nil)
	userService := services.NewUserService(mockUserRepository)
	err := userService.CreateUser(&models.User{})
	if err != nil {
		t.Errorf("Error creating user: %s", err)
	}
}

func testGetAllUsers(t *testing.T, mock *gomock.Controller) {
	mockUserRepository := mocks.NewMockUserRepository(mock)
	mockUserRepository.EXPECT().GetAll().Return([]models.User{
		{
			ID: uuid.New(),
		},
		{
			ID: uuid.New(),
		},
	})
	userService := services.NewUserService(mockUserRepository)
	users := userService.GetAllUsers()
	if len(users) != 2 {
		t.Errorf("Expected 2 users to be returned")
	}
}

func testGetUserById_UserNotExist(t *testing.T, mock *gomock.Controller) {
	var notExistingUserID = uuid.New()
	mockUserRepository := mocks.NewMockUserRepository(mock)
	mockUserRepository.EXPECT().Get(notExistingUserID).Return(nil, errors.New("user not found"))
	userService := services.NewUserService(mockUserRepository)
	user, err := userService.GetUserById(notExistingUserID)
	if user != nil {
		t.Errorf("Expected user to be nil")
	}
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func testGetUserById_UserExist(t *testing.T, mock *gomock.Controller) {
	var existingUserID = uuid.New()
	mockUserRepository := mocks.NewMockUserRepository(mock)
	mockUserRepository.EXPECT().Get(existingUserID).Return(&models.User{
		ID: existingUserID,
	}, nil)
	userService := services.NewUserService(mockUserRepository)
	user, err := userService.GetUserById(existingUserID)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}
	if user.ID != existingUserID {
		t.Errorf("User ID is not correct")
	}
}
