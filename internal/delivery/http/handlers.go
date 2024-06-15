package http

import (
	"github.com/bagasunix/go-clean-architecture/internal/controllers"
	"github.com/bagasunix/go-clean-architecture/internal/delivery/http/handlers"
	"github.com/bagasunix/go-clean-architecture/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App               *fiber.App
	UserController    controllers.UserController
	WelcomeController controllers.WelcomeController
	Cfg               *config.Cfg
}

func NewHttpHandler(r RouteConfig) *fiber.App {
	// Handlers
	handlers.MakeUserHandler(r.UserController, r.App.Group(r.Cfg.Server.Version+"/user").(*fiber.Group))
	handlers.MakeWelcomeHandler(r.WelcomeController, r.App.Group(r.Cfg.Server.Version+"").(*fiber.Group))
	return r.App
}
