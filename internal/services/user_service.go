package services

import (
	"context"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"mai_lab/internal/apperror"
	"mai_lab/internal/domain/models"
	"mai_lab/internal/infrastructure/storage"
)

type UserService interface {
	CreateUser(ctx context.Context, dto models.CreateUserDTO) error
	GetUserByID(ctx context.Context, id uuid.UUID) (models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, dto models.UpdateUserDTO) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type service struct {
	storage storage.Storage
}

func NewService(userStorage storage.Storage) UserService {
	return &service{storage: userStorage}
}

func (s *service) CreateUser(ctx context.Context, dto models.CreateUserDTO) error {
	u := dto.NewUser()
	if err := s.storage.Create(ctx, u); err != nil {
		return err
	}
	return nil
}

func (s *service) GetUserByID(ctx context.Context, id uuid.UUID) (models.User, error) {
	u, err := s.storage.GetUser(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (s *service) GetAllUsers(ctx context.Context) ([]models.User, error) {
	all, err := s.storage.GetAll(ctx)
	if err != nil {
		return []models.User{}, err
	}
	return all, nil
}

func (s *service) UpdateUser(ctx context.Context, dto models.UpdateUserDTO) error {

	user, err := s.storage.GetUser(ctx, dto.ID)
	if err != nil {
		return err
	}

	log.Println("compare old and new passwords")
	if dto.OldPassword != dto.NewPassword || dto.NewPassword == "" {
		log.Println("compare hash current password and old password")
		if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(dto.OldPassword)); err != nil {
			return apperror.BadRequestError("old password does not match current password")
		}
		dto.Password = dto.NewPassword
	}

	dto.UpdateUser(&user)

	if err := s.storage.Update(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if err := s.storage.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
