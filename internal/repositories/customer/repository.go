package customer

import (
	"context"

	"github.com/bagasunix/go-clean-architecture/internal/entities"
	"github.com/bagasunix/go-clean-architecture/internal/models"
	"github.com/gofrs/uuid"
)

type Command interface {
	Create(ctx context.Context, model *entities.Customer) error
	Update(ctx context.Context, model *entities.Customer) error
	Delete(ctx context.Context, id uuid.UUID) error
}
type Query interface {
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*entities.Customer])
	GetByEmail(ctx context.Context, email string) (result models.SingleResult[*entities.Customer])
	GetByPhone(ctx context.Context, phone string) (result models.SingleResult[*entities.Customer])
}

type Repository interface {
	Command
	Query
}
