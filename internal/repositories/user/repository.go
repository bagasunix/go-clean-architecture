package user

import (
	"context"

	"github.com/bagasunix/go-clean-architecture/internal/entities"
	"github.com/bagasunix/go-clean-architecture/internal/models"
	"github.com/gofrs/uuid"
)

type Command interface {
	Create(ctx context.Context, model *entities.User) error
	Update(ctx context.Context, maps *map[string]any, id uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Query interface {
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*entities.User])
	GetByEmail(ctx context.Context, email string) (result models.SingleResult[*entities.User])
}

type Repository interface {
	Command
	Query
}
