package services

import (
	"context"
	"log"

	"github.com/google/uuid"
	"mai_lab/internal/domain/models"
	"mai_lab/internal/infrastructure/storage"
	"mai_lab/rpc"
)

type TwirpService struct {
	storage storage.Storage
}

func NewTwirpService(userStorage storage.Storage) rpc.Users {
	return &TwirpService{storage: userStorage}

}

// CreateUser creates new user
func (s *TwirpService) CreateUser(ctx context.Context, r *rpc.CreateUserRequest) (*rpc.CreateUserResponse, error) {
	userDTO := models.CreateUserDTO{
		Name:     r.Name,
		Email:    r.Email,
		Mobile:   r.Mobile,
		Password: r.Password,
	}
	user := userDTO.NewUser()
	if err := s.storage.Create(ctx, user); err != nil {
		return nil, err
	}
	response := models.TwirpFromUser(&user)
	return &rpc.CreateUserResponse{User: response}, nil
}

func (s *TwirpService) GetUser(ctx context.Context, r *rpc.GetUserRequest) (*rpc.GetUserResponse, error) {
	id, err := uuid.Parse(r.Id)
	if err != nil {
		return nil, err
	}

	user, err := s.storage.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	response := models.TwirpFromUser(&user)
	return &rpc.GetUserResponse{User: response}, nil
}

func (s *TwirpService) GetAllUsers(ctx context.Context, r *rpc.GetAllUsersRequest) (*rpc.GetAllUsersResponse, error) {
	users, err := s.storage.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	twirpUsers := models.TwirpFromUsers(users)

	return &rpc.GetAllUsersResponse{Users: twirpUsers}, nil

}

func (s *TwirpService) UpdateUser(ctx context.Context, r *rpc.UpdateUserRequest) (*rpc.UpdateUserResponse, error) {
	log.Fatal("Implement me")
	return nil, nil
}

func (s *TwirpService) DeleteUser(ctx context.Context, r *rpc.DeleteUserRequest) (*rpc.DeleteUserResponse, error) {
	log.Fatal("Implement me")
	return nil, nil
}
