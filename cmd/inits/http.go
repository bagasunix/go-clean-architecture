package inits

import (
	"github.com/bagasunix/go-clean-architecture/internal/delivery/http"
	transportHttp "github.com/bagasunix/go-clean-architecture/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

func InitHttpHandler(f *http.RouteConfig) *fiber.App {
	return transportHttp.NewHttpHandler(*f)
}
