package storage

import (
	"context"
	"github.com/google/uuid"
	"mai_lab/internal/domain/models"
)

type Storage interface {
	Create(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, id uuid.UUID) (models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
