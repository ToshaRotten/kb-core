package main

import (
	"log"
	"log/slog"
	"main/config"
	"main/storage"

	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg := config.MustLoadConfig("configs/local.env")
	app := fiber.New()
	storage, err := storage.New()
	if err != nil {
		slog.Error(err.Error())
	}
	app.Use(func(c fiber.Ctx) error {
		slog.Info(c.Request().URI())
		return c.Next()
	})

	app.Get("/professor", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/group", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/lesson", func(c fiber.Ctx) error {
		c.JSON()
		return c.SendString("Hello World!")
	})

	log.Fatal(app.Listen(cfg.HTTPServer.Host))
}
