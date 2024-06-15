package inits

import (
	"github.com/bagasunix/go-clean-architecture/internal/controllers"
	"github.com/bagasunix/go-clean-architecture/internal/delivery/http"
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

func SetupApp(app *SetupAppConfig) *http.RouteConfig {
	// setup repositories
	userRepository := repositories.New(app.Log, app.DB)
	// setup use cases
	userUseCase := usecases.NewUser(app.Log, userRepository)
	// setup controller
	userController := controllers.NewUserController(userUseCase, app.Cfg, app.Log)
	welcomeController := controllers.NewWelcomeController(app.Cfg)

	return &http.RouteConfig{
		App:               app.App,
		UserController:    userController,
		WelcomeController: welcomeController,
		Cfg:               app.Cfg,
	}
}
