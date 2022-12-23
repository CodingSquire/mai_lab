// Package controllers contains all the controllers for the application
package controllers

import (
	"encoding/json"
	"net/http"
	"users/internal/api/common"
	"users/internal/api/common/dtos"
	"users/internal/contracts"

	"github.com/google/uuid"
)

// UserController is an interface for user controllers.
type UserController interface {
	GetUserById(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	service contracts.UserService
}

// NewUserController returns a new instance of UserController.
func NewUserController(service contracts.UserService) UserController {
	return &userController{
		service: service,
	}
}

func prepareResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

// GetUserById returns a user by id provided in context.
// Returns a 404 if the user does not exist.
// Returns a 400 if the id is not a valid uuid.
// Returns a 200 with the user if the user exists.
func (c *userController) GetUserById(w http.ResponseWriter, r *http.Request) {
	prepareResponse(w, r)
	id := r.Context().Value(common.ContextKeyParams).(map[string]string)["id"]
	userId, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := c.service.GetUserById(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userResponse := dtos.UserResponseDto{}
	userResponse.FromUser(user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userResponse)
}

// GetAllUsers returns all users.
// Returns a 200 with the users.
func (c *userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	prepareResponse(w, r)
	users := c.service.GetAllUsers()

	usersResponse := []dtos.UserResponseDto{}
	for _, user := range users {
		userResponse := dtos.UserResponseDto{}

		usersResponse = append(usersResponse, userResponse.FromUser(&user))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usersResponse)
}

// CreateUser creates a new user.
// Returns a 400 if the request body is invalid.
// Returns a 400 if the user already exists.
// Returns a 201 with the created user.
func (c *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	prepareResponse(w, r)
	var userRequest dtos.UserRequestDto
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userValidator := dtos.NewUserValidator(userRequest)
	if !userValidator.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(userValidator.Errors)
		return
	}

	user := userRequest.ToUser()

	err = c.service.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	userResponse := dtos.UserResponseDto{}
	userResponse.FromUser(user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)
}

// UpdateUser updates a user.
// Returns a 400 if the request body is invalid.
// Returns a 404 if the user does not exist.
// Returns a 400 if the id is not a valid uuid.
// Returns a 200 with the updated user.
func (c *userController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	prepareResponse(w, r)
	id := r.Context().Value(common.ContextKeyParams).(map[string]string)["id"]
	userId, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userRequest dtos.UserRequestDto
	err = json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userValidator := dtos.NewUserValidator(userRequest)
	if !userValidator.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(userValidator.Errors)
		return
	}

	user := userRequest.ToUser()
	user.ID = userId

	err = c.service.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	userResponse := dtos.UserResponseDto{}
	userResponse.FromUser(user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userResponse)
}

// DeleteUser deletes a user.
// Returns a 404 if the user does not exist.
// Returns a 400 if the id is not a valid uuid.
// Returns a 200 if the user was deleted.
func (c *userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	prepareResponse(w, r)
	id := r.Context().Value(common.ContextKeyParams).(map[string]string)["id"]
	userId, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.service.DeleteUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
