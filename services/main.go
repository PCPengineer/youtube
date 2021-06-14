package main

import (
	"github.com/gofiber/fiber/v2"
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
	if err := app.Listen(":7575"); err != nil {
		panic(err)
	}
}

func setUpRoute(app *fiber.App) {
	//youtube := app.Group("/youtube")
}