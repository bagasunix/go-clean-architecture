package repositories

import (
	"github.com/bagasunix/go-clean-architecture/internal/repositories/address"
	"github.com/bagasunix/go-clean-architecture/internal/repositories/customer"
	"github.com/bagasunix/go-clean-architecture/internal/repositories/order"
	"github.com/bagasunix/go-clean-architecture/internal/repositories/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repositories interface {
	GetUser() user.Repository
	GetOrder() order.Repository
	GetCustomer() customer.Repository
	GetAddressCustomer() address.Repository
}

type repo struct {
	user             user.Repository
	order            order.Repository
	customer         customer.Repository
	address_customer address.Repository
}

// GetAddressCustomer implements Repositories.
func (r *repo) GetAddressCustomer() address.Repository {
	return r.address_customer
}

// GetCustomer implements Repositories.
func (r *repo) GetCustomer() customer.Repository {
	return r.customer
}

// GetOrder implements Repositories.
func (r *repo) GetOrder() order.Repository {
	return r.order
}

// GetUser implements Repositories.
func (r *repo) GetUser() user.Repository {
	return r.user
}

func New(logger *zap.Logger, db *gorm.DB) Repositories {
	rs := new(repo)
	rs.customer = customer.NewGorm(logger, db)
	rs.order = order.NewGorm(logger, db)
	rs.address_customer = address.NewGorm(logger, db)
	rs.user = user.NewGorm(logger, db)
	return rs
}
