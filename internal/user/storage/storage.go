package storage

import (
	"context"
	"github.com/google/uuid"
	"mai_lab/internal/user/model"
)

type Storage interface {
	Create(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, id string) (model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
