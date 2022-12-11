package user

import (
	"context"
	"github.com/google/uuid"
	"mai_lab/pkg"
)

type Service interface {
	CreateUser(ctx context.Context, dto CreateUserDTO) error
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	UpdateUser(ctx context.Context, dto CreateUserDTO) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type service struct {
	storage Storage
}

func NewService(s Storage) Service {
	return &service{storage: s}
}

func (s *service) CreateUser(ctx context.Context, dto CreateUserDTO) error {
	u := User{
		ID:           uuid.New(),
		Name:         dto.Name,
		Email:        dto.Email,
		Mobile:       dto.Mobile,
		PasswordHash: pkg.GeneratePasswordHash(dto.Password),
	}

	if err := s.storage.Create(ctx, u); err != nil {
		return err
	}
	return nil
}

func (s *service) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	u, err := s.storage.GetUser(ctx, id)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (s *service) UpdateUser(ctx context.Context, dto CreateUserDTO) error {
	// TODO rewrite the user update logic
	u := User{
		ID:           uuid.New(),
		Name:         dto.Name,
		Email:        dto.Email,
		Mobile:       dto.Mobile,
		PasswordHash: pkg.GeneratePasswordHash(dto.Password),
	}

	if err := s.storage.Update(ctx, u); err != nil {
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
