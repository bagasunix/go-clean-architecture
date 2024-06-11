package address

import (
	"context"

	"github.com/bagasunix/go-clean-architecture/internal/entities"
	"github.com/gofrs/uuid"
)

type Command interface {
	Create(ctx context.Context, model *entities.AddressCustomer) error
	Update(ctx context.Context, model *entities.AddressCustomer) error
	Delete(ctx context.Context, id uuid.UUID) error
}
type Query interface {
}

type Repository interface {
	Command
	Query
}
