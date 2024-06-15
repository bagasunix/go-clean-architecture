package controllers

import (
	"github.com/bagasunix/go-clean-architecture/internal/models"
	"github.com/bagasunix/go-clean-architecture/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type welcomeController struct {
	cfg *config.Cfg
}

// Welcome implements WelcomeController.
func (w *welcomeController) Welcome(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusOK).JSON(models.BaseResponse[any]{
		Message: "Welcome to API Go Clean version " + w.cfg.General.AppVersion + ", enjoy and chersss :)",
	})
	return nil
}

type WelcomeController interface {
	Welcome(ctx *fiber.Ctx) error
}

func NewWelcomeController(cfg *config.Cfg) WelcomeController {
	c := new(welcomeController)
	c.cfg = cfg
	return c
}
