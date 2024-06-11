package configs

import (
	"github.com/bagasunix/go-clean-architecture/internal/models"
	"github.com/bagasunix/go-clean-architecture/pkg/config"
	"github.com/bagasunix/go-clean-architecture/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

func NewFiber(config *config.Cfg, logger *zap.Logger) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:           config.General.AppName,
		Prefork:           config.Server.Prefork,
		UnescapePath:      true,
		ErrorHandler:      NewErrorHandler(logger),
		EnablePrintRoutes: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Security-Policy", "default-src 'self'")
		return c.Next()
	})
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(favicon.New())

	return app
}

func NewErrorHandler(logger *zap.Logger) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		errors.HandlerWithLoggerReturnedError(logger, err)
		e := err.(*fiber.Error)
		ctx.Status(e.Code)
		ctx.Set("Content-Type", "application/json; charset=utf-8")
		responseBody := new(models.BaseResponse[any])
		responseBody.Data = e.Error()
		responseBody.Errors = e.Message
		return ctx.JSON(responseBody)
	}
}
