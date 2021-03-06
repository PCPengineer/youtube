package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"youtube/pkg/errorHandler"
	"youtube/pkg/logger"
)

func main() {
	logger.Init()
	app := fiber.New(fiber.Config{
		ErrorHandler:          errorHandler.DefaultErrHandler,
		DisableStartupMessage: true,
	})
	setUpRoute(app)
	if err := errorHandler.LoadLocale(); err != nil {
		panic(err)
	}
	zap.L().Info("start listening on port 7575")
	if err := app.Listen(":7575"); err != nil {
		panic(err)
	}
}
