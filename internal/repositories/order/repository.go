package order

import (
	"context"

	"github.com/bagasunix/go-clean-architecture/internal/entities"
	"github.com/bagasunix/go-clean-architecture/internal/models"
	"github.com/gofrs/uuid"
)

type Command interface {
	Create(ctx context.Context, model *entities.Order) error
	Update(ctx context.Context, model *entities.Order) error
	Delete(ctx context.Context, id uuid.UUID) error
}
type Query interface {
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*entities.Order])
}

type Repository interface {
	Command
	Query
}
