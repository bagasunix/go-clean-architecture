package inits

import (
	"github.com/bagasunix/go-clean-architecture/internal/controllers"
	"github.com/bagasunix/go-clean-architecture/internal/delivery/http"
	"github.com/bagasunix/go-clean-architecture/internal/delivery/http/middlewares"
	"github.com/bagasunix/go-clean-architecture/internal/repositories"
	"github.com/bagasunix/go-clean-architecture/internal/usecases"
	"github.com/bagasunix/go-clean-architecture/pkg/config"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SetupAppConfig struct {
	DB  *gorm.DB
	App *fiber.App
	Log *zap.Logger
	Cfg *config.Cfg
}

// InitService menginisialisasi service dengan logger dan repositories, serta menetapkan middlewares
func InitUsecase(logger *zap.Logger, repo repositories.Repositories) usecases.UsecasesContract {
	serviceBuilder := usecases.NewUsecaseBuilder(logger, repo)
	middlewares := []usecases.Middleware{
		middlewares.LoggingMiddleware(logger, repo),
	}
	serviceBuilder.SetMiddlewares(middlewares)
	return serviceBuilder.Build()
}

func SetupApp(app *SetupAppConfig) *http.RouteConfig {
	// setup repositories
	// userRepository := repositories.New(app.Log, app.DB)
	repository := repositories.New(app.Log, app.DB)
	// setup use cases
	// userUseCase := usecases.NewUser(app.Log, userRepository)
	usecase := InitUsecase(app.Log, repository)
	// setup controller
	userController := controllers.NewUserController(usecase, app.Cfg, app.Log)
	welcomeController := controllers.NewWelcomeController(app.Cfg)

	return &http.RouteConfig{
		App:               app.App,
		UserController:    userController,
		WelcomeController: welcomeController,
		Cfg:               app.Cfg,
	}
}
