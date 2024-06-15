package models

import (
	"encoding/json"

	"github.com/bagasunix/go-clean-architecture/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type BaseResponse[T any] struct {
	Code    int           `json:"code,omitempty"`
	Message string        `json:"message"`
	Data    *T            `json:"data,omitempty"`
	Paging  *PageMetadata `json:"paging,omitempty"`
	Errors  error         `json:"errors,omitempty"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}

// WriteResponse function
func WriteResponse[T any](ctx *fiber.Ctx, cfg *config.Cfg, result any, err error) error {
	resp := new(BaseResponse[T])
	// Mengatur header
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	// Marshal data menjadi JSON
	bytes, _ := json.Marshal(result)
	// Unmarshal JSON ke struct CustomError
	json.Unmarshal(bytes, &resp)
	if err != nil {
		if cfg.Server.Env != "dev" {
			resp.Data = nil
		}
		resp.Errors = err
		return ctx.Status(resp.Code).JSON(resp)
	} else {
		return ctx.Status(resp.Code).JSON(resp)
	}
}
