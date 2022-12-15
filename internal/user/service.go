package user

import (
	"context"
	"github.com/google/uuid"
	"log"
	"mai_lab/internal/user/model"
	storage2 "mai_lab/internal/user/storage"
	"mai_lab/pkg"
)

type Service interface {
	CreateUser(ctx context.Context, dto model.CreateUserDTO) error
	GetUserByID(ctx context.Context, id string) (model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, dto model.UpdateUserDTO) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type service struct {
	storage storage2.Storage
}

func NewService(userStorage storage2.Storage) Service {
	return &service{storage: userStorage}
}

func (s *service) CreateUser(ctx context.Context, dto model.CreateUserDTO) error {
	u := dto.NewUser()
	if err := s.storage.Create(ctx, u); err != nil {
		return err
	}
	return nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (model.User, error) {
	u, err := s.storage.GetUser(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (s *service) GetAllUsers(ctx context.Context) ([]model.User, error) {
	all, err := s.storage.GetAll(ctx)
	if err != nil {
		return []model.User{}, err
	}
	return all, nil
}

func (s *service) UpdateUser(ctx context.Context, dto model.UpdateUserDTO) error {
	var updatedUser model.User
	log.Println("compare old and new passwords")
	if dto.OldPassword != dto.NewPassword || dto.NewPassword == "" {

		//TODO implement
		//user, err := s.storage.GetUser(ctx, dto.ID)
		//if err != nil {
		//	return err
		//}
		//log.Println("compare hash current password and old password")
		//if pkg.GeneratePasswordHash(dto.OldPassword) != user.PasswordHash {
		//	return apperror.BadRequestError("old password does not match current password")
		//}
		//dto.Password = dto.NewPassword
	}
	updatedUser = dto.UpdateUser()

	log.Println("generate password hash")
	updatedUser.PasswordHash = pkg.GeneratePasswordHash(dto.NewPassword)

	if err := s.storage.Update(ctx, updatedUser); err != nil {
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
