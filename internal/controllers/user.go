package controllers

import (
	"github.com/bagasunix/go-clean-architecture/internal/models"
	"github.com/bagasunix/go-clean-architecture/internal/usecases"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserController interface {
	CreateUser(ctx *fiber.Ctx) error
}

type userController struct {
	Log     *zap.Logger
	useCase usecases.UserEndpoint
}

// CreateUser implements UserController.
func (u *userController) CreateUser(ctx *fiber.Ctx) error {
	request := new(models.CreateUser)
	if err := ctx.BodyParser(&request); err != nil {
		u.Log.Error("failed to parse request body", zap.Error(err))
		return fiber.ErrBadRequest
	}
	resp, err := u.useCase.CreateUser(ctx, request)
	if err != nil {
		u.Log.Error(err.Error())
		return err
	}
	return ctx.Status(201).JSON(models.BaseResponse[*models.ResponseUser]{Data: &resp.Data, Message: resp.Message})
}

func NewUserController(useCase usecases.UserEndpoint, logger *zap.Logger) UserController {
	c := new(userController)
	c.useCase = useCase
	c.Log = logger
	return c
}
