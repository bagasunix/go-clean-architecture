package inits

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bagasunix/go-clean-architecture/pkg/config"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := InitLogger()
	defer logger.Sync()
	config := config.InitConfig(ctx, logger)
	db := InitDb(ctx, logger, config)
	// repositories := repositories.New(logger, db)

	app := NewFiber(config, logger)
	setup := SetupApp(&SetupAppConfig{
		DB:  db,
		App: app,
		Log: logger,
		Cfg: config,
	})
	httpHandler := InitHttpHandler(setup)

	errs := make(chan error)
	defer close(errs)
	go initCancel(errs)
	go initHttp(httpHandler, config, errs)
	logger.Error("exit", zap.Error(<-errs))
}

func initCancel(errs chan error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	errs <- fmt.Errorf("%s", <-c)
}

func initHttp(c *fiber.App, config *config.Cfg, errs chan error) {
	errs <- c.Listen(":" + config.Server.Port)
}
