package handlers

import (
	"github.com/bagasunix/go-clean-architecture/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func MakeUserHandler(endpoints controllers.UserController, r *fiber.Group) *fiber.Group {
	r.Post("", endpoints.CreateUser)
	return r
}
