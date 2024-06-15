package handlers

import (
	"github.com/bagasunix/go-clean-architecture/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func MakeWelcomeHandler(endpoints controllers.WelcomeController, r *fiber.Group) *fiber.Group {
	r.Get("", endpoints.Welcome)
	return r
}
