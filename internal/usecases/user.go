package usecases

import (
	"strconv"
	"strings"

	"github.com/bagasunix/go-clean-architecture/internal/entities"
	"github.com/bagasunix/go-clean-architecture/internal/models"
	"github.com/bagasunix/go-clean-architecture/internal/repositories"
	"github.com/bagasunix/go-clean-architecture/pkg/config"
	"github.com/bagasunix/go-clean-architecture/pkg/errors"
	"github.com/bagasunix/go-clean-architecture/pkg/helpers"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserEndpoint interface {
	CreateUser(ctx *fiber.Ctx, req *models.CreateUser) (response *models.BaseResponse[models.ResponseUser], err error)
}
type userUseCase struct {
	repo   repositories.Repositories
	logger *zap.Logger
	Cfg    *config.Cfg
}

// CreateUser implements UserEndpoint.
func (u *userUseCase) CreateUser(ctx *fiber.Ctx, req *models.CreateUser) (response *models.BaseResponse[models.ResponseUser], err error) {
	responseBuilder := new(models.BaseResponse[models.ResponseUser])
	if req.Validate() != nil {
		responseBuilder.Code = fiber.StatusBadRequest
		responseBuilder.Message = "Validasi error"
		return responseBuilder, req.Validate()
	}

	checkEmail := u.repo.GetUser().GetByEmail(ctx.Context(), req.Email)
	if checkEmail.Value.Email == req.Email {
		responseBuilder.Code = fiber.StatusConflict
		responseBuilder.Message = "Email sudah terdaftar"
		return responseBuilder, errors.CustomError("email " + errors.ERR_ALREADY_EXISTS)
	}
	if checkEmail.Error != nil {
		responseBuilder.Code = fiber.StatusConflict
		responseBuilder.Message = "Validasi email invalid"
		return responseBuilder, checkEmail.Error
	}

	switch strings.ToLower(req.Sex) {
	case "laki-laki":
		req.Sex = "1"
	case "perempuan":
		req.Sex = "0"
	default:
		responseBuilder.Code = fiber.StatusConflict
		responseBuilder.Message = "Jenis kelamin tidak valid"
		return responseBuilder, errors.ErrInvalidAttributes("sex")
	}

	intSex, _ := strconv.Atoi(req.Sex)
	entityBuild := new(entities.User)
	entityBuild.ID = helpers.GenerateUUIDV4(u.logger)
	entityBuild.FullName = req.FullName
	entityBuild.Username = req.Username
	entityBuild.Sex = intSex
	entityBuild.Email = req.Email
	entityBuild.Password = helpers.HashAndSalt([]byte(req.Password))
	entityBuild.IsActive = 1

	if err = u.repo.GetUser().Create(ctx.Context(), entityBuild); err != nil {
		responseBuilder.Code = fiber.StatusConflict
		responseBuilder.Message = "Gagal membuat pengguna"
		return responseBuilder, err
	}

	mBuild := new(models.ResponseUser)
	mBuild.ID = entityBuild.ID
	mBuild.FullName = entityBuild.FullName
	mBuild.Email = entityBuild.Email
	switch entityBuild.Sex {
	case 1:
		mBuild.Sex = "Laki-laki"
	case 0:
		mBuild.Sex = "Perempuan"
	}
	mBuild.IsActive = entityBuild.IsActive
	mBuild.CreatedAt = entityBuild.CreatedAt

	responseBuilder.Code = fiber.StatusCreated
	responseBuilder.Message = "User berhasil dibuat"
	responseBuilder.Data = mBuild
	return responseBuilder, err
}

func NewUser(logger *zap.Logger, repo repositories.Repositories) UserEndpoint {
	a := new(userUseCase)
	a.repo = repo
	a.logger = logger
	return a
}
