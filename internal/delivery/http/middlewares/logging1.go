package middlewares

import (
	"time"

	"github.com/bagasunix/go-clean-architecture/internal/models"
	"github.com/bagasunix/go-clean-architecture/internal/repositories"
	"github.com/bagasunix/go-clean-architecture/internal/usecases"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type loggingMiddleware struct {
	logger *zap.Logger
	next   usecases.UsecasesContract
}

// CreateUser implements internal.Service.
func (l *loggingMiddleware) CreateUser(ctx *fiber.Ctx, req *models.CreateUser) (response *models.BaseResponse[models.ResponseUser], err error) {
	defer func(begin time.Time) {
		l.logger.Info("Create User", zap.String("request", string(req.ToJSON())), zap.String("response", string(response.ToJSON())), zap.Error(err), zap.Duration("took", time.Since(begin)))
	}(time.Now())
	return l.next.CreateUser(ctx, req)
}

func LoggingMiddleware(logger *zap.Logger, repo repositories.Repositories) usecases.Middleware {
	return func(repo repositories.Repositories, next usecases.UsecasesContract) usecases.UsecasesContract {
		return &loggingMiddleware{logger: logger, next: next}
	}
}
