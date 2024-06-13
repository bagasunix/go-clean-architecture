package models

import (
	"github.com/bagasunix/go-clean-architecture/pkg/config"
	"github.com/gofiber/fiber/v2"
)

var cfg *config.Cfg

type BaseResponse[T any] struct {
	Message string        `json:"message"`
	Data    *T            `json:"data"`
	Paging  *PageMetadata `json:"paging,omitempty"`
	Errors  string        `json:"errors,omitempty"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}

func WriteResponse[T any](ctx *fiber.Ctx, resp BaseResponse[T], statusCode int) error {
	if resp.Errors != "" {
		if cfg.Server.Env != "dev" {
			resp.Data = nil
		}
		return ctx.Status(statusCode).JSON(resp)
	} else {
		return ctx.Status(statusCode).JSON(resp)
	}
}
