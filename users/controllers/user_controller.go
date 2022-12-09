package controllers

import (
	"encoding/json"
	"net/http"
	"users/ctxkeys"
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

func prepareResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (c *userController) GetUserById(w http.ResponseWriter, r *http.Request) {
	prepareResponse(w, r)
	id := r.Context().Value(ctxkeys.ContextKeyParams).(map[string]string)["id"]
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

func (c *userController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	prepareResponse(w, r)
	id := r.Context().Value(ctxkeys.ContextKeyParams).(map[string]string)["id"]
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

func (c *userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	prepareResponse(w, r)
	id := r.Context().Value(ctxkeys.ContextKeyParams).(map[string]string)["id"]
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
