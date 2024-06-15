package customer

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

// GetByPhone implements Repository.
func (g *gormProvider) GetByPhone(ctx context.Context, phone string) (result models.SingleResult[*entities.Customer]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Where("phone = ?", phone).Find(&result.Value).Error)
	return result
}

// GetByEmail implements Repository.
func (g *gormProvider) GetByEmail(ctx context.Context, email string) (result models.SingleResult[*entities.Customer]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Where("email = ?", email).Find(&result.Value).Error)
	return result
}

// Create implements Repository.
func (g *gormProvider) Create(ctx context.Context, model *entities.Customer) error {
	return errors.ErrDuplicateValue(g.logger, g.GetModelName(), g.db.WithContext(ctx).Create(&model).Error)
}

// Delete implements Repository.
func (g *gormProvider) Delete(ctx context.Context, id uuid.UUID) error {
	return errors.ErrSomethingWrong(g.logger, g.db.WithContext(ctx).Delete(&entities.Customer{}, "id = ?", id).Error)
}

// GetById implements Repository.
func (g *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*entities.Customer]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Where("id = ?", id).First(&result.Value).Error)
	return result
}

// GetConnection implements Repository.
func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

// GetModelName implements Repository.
func (g *gormProvider) GetModelName() string {
	return "customer"
}

// Update implements Repository.
func (g *gormProvider) Update(ctx context.Context, model *entities.Customer) error {
	return errors.ErrDuplicateValue(g.logger, g.GetModelName(), g.db.WithContext(ctx).Updates(model).Error)
}
func NewGorm(logger *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.db = db
	g.logger = logger
	return g
}
