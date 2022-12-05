package controllers

import (
	"encoding/json"
	"net/http"
	"users/dtos"
	"users/services"

	"github.com/google/uuid"
)

type UserController interface {
	GetUserById(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResponse)
}

func (c *userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := c.service.GetAllUsers()

	var usersResponse []dtos.UserResponseDto
	for _, user := range users {
		userResponse := dtos.UserResponseDto{}
		usersResponse = append(usersResponse, userResponse.FromUser(&user))
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usersResponse)
}

func (c *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest dtos.UserRequestDto
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userValidator := dtos.NewUserValidator(userRequest)
	if !userValidator.IsValid() {
		json.NewEncoder(w).Encode(userValidator.Errors)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := userRequest.ToUser()

	err = c.service.CreateUser(user)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *userController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
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
		json.NewEncoder(w).Encode(userValidator.Errors)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := userRequest.ToUser()
	user.ID = userId

	err = c.service.UpdateUser(user)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.service.DeleteUser(userId)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
