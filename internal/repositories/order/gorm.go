package order

import (
	"context"

	"github.com/bagasunix/go-clean-architecture/internal/entities"
	"github.com/bagasunix/go-clean-architecture/internal/models"
	"github.com/bagasunix/go-clean-architecture/pkg/errors"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type gormProvider struct {
	db     *gorm.DB
	logger *zap.Logger
}

// Create implements Repository.
func (g *gormProvider) Create(ctx context.Context, model *entities.Order) error {
	return errors.ErrDuplicateValue(g.logger, g.GetModelName(), g.db.WithContext(ctx).Create(&model).Error)
}

// Delete implements Repository.
func (g *gormProvider) Delete(ctx context.Context, id uuid.UUID) error {
	return errors.ErrSomethingWrong(g.logger, g.db.WithContext(ctx).Delete(&entities.Order{}, "id = ?", id).Error)
}

// GetById implements Repository.
func (g *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*entities.Order]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Where("id = ?", id).First(&result.Value).Error)
	return result
}

// GetConnection implements Repository.
func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

// GetModelName implements Repository.
func (g *gormProvider) GetModelName() string {
	return "order"
}

// Update implements Repository.
func (g *gormProvider) Update(ctx context.Context, model *entities.Order) error {
	return errors.ErrDuplicateValue(g.logger, g.GetModelName(), g.db.WithContext(ctx).Updates(model).Error)
}
func NewGorm(logger *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.db = db
	g.logger = logger
	return g
}
