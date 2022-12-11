package user

import (
	"context"
	"github.com/google/uuid"
)

type Storage interface {
	Create(ctx context.Context, user User) error
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, user User) error
}
