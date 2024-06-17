package controllers

import (
	"github.com/bagasunix/go-clean-architecture/internal/models"
	"github.com/bagasunix/go-clean-architecture/internal/usecases"
	"github.com/bagasunix/go-clean-architecture/pkg/config"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserController interface {
	CreateUser(ctx *fiber.Ctx) error
}

type userController struct {
	Log     *zap.Logger
	useCase usecases.UserEndpoint
	Cfg     *config.Cfg
}

// CreateUser implements UserController.
func (u *userController) CreateUser(ctx *fiber.Ctx) error {
	request := new(models.CreateUser)
	if err := ctx.BodyParser(&request); err != nil {
		// u.Log.Error("failed to parse request body", zap.Error(err))
		return fiber.ErrBadRequest
	}
	resp, err := u.useCase.CreateUser(ctx, request)
	if err != nil {
		// u.Log.Error(err.Error())
		return models.WriteResponse[models.ResponseUser](ctx, u.Cfg, resp, err)
	}
	return models.WriteResponse[models.ResponseUser](ctx, u.Cfg, resp, nil)
}

func NewUserController(useCase usecases.UserEndpoint, cfg *config.Cfg, logger *zap.Logger) UserController {
	c := new(userController)
	c.useCase = useCase
	c.Log = logger
	c.Cfg = cfg
	return c
}
