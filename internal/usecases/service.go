package usecases

import (
	"github.com/bagasunix/go-clean-architecture/internal/repositories"
	"go.uber.org/zap"
)

// Usecases interface mencakup metode dari usecase
type UsecasesContract interface {
	UserEndpoint
}

// struct service mengimplementasikan Service
type usecases struct {
	UserEndpoint
}

// Middleware adalah function type yang menerima repositories dan Usecases, mengembalikan Usecases
type Middleware func(repo repositories.Repositories, contract UsecasesContract) UsecasesContract

// UsecaseBuilder membantu membangun Usecases dengan logger, repositories, dan middlewares
type UsecaseBuilder struct {
	logger       *zap.Logger
	repositories repositories.Repositories
	middlewares  []Middleware
}

// NewUsecaseBuilder membuat instance baru dari UsecaseBuilder
func NewUsecaseBuilder(logger *zap.Logger, repositories repositories.Repositories) *UsecaseBuilder {
	return &UsecaseBuilder{
		logger:       logger,
		repositories: repositories,
	}
}

// buildUsecases membuat dan menginisialisasi Usecases baru
func buildUsecases(logger *zap.Logger, repo repositories.Repositories) UsecasesContract {
	return &usecases{
		UserEndpoint: NewUser(logger, repo),
	}
}

// SetMiddlewares menetapkan middlewares ke UsecaseBuilder
func (s *UsecaseBuilder) SetMiddlewares(middlewares []Middleware) *UsecaseBuilder {
	s.middlewares = middlewares
	return s
}

// Build membangun Usecases dengan middlewares
func (s *UsecaseBuilder) Build() UsecasesContract {
	svc := buildUsecases(s.logger, s.repositories)
	for _, m := range s.middlewares {
		svc = m(s.repositories, svc)
	}
	return svc
}
