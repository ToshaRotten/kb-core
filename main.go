package main

import (
	"log"
	"log/slog"
	"main/config"
	"main/models/database"
	"main/storage"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.MustLoadConfig("configs/local.env")
	app := fiber.New()

	app.Use(cors.New())

	storage, err := storage.New(cfg.Storage)
	if err != nil {
		slog.Error(err.Error())
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: storage.DB,
	}), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
	}

	err = gormDB.AutoMigrate(
		&database.Professor{},
		&database.User{},
		&database.Student{},
		&database.Group{},
		&database.Lesson{},
	)
	if err != nil {
		slog.Error(err.Error())
	}

	app.Use(func(c fiber.Ctx) error {
		slog.Info(string(c.Request().RequestURI()))
		return c.Next()
	})

	//
	// app.Get("/professor/:id", func(c fiber.Ctx) error {
	// 	id := c.Params("id", "0")
	// 	professor, err := storage.GetProfessor(id)
	// 	if err != nil {
	// 		return fiber.ErrNotFound
	// 	}
	//
	// 	return c.JSON(professor)
	// })
	//
	// app.Get("/group", func(c fiber.Ctx) error {
	// 	return c.SendString("Hello, World 👋!")
	// })
	//
	// app.Get("/lesson", func(c fiber.Ctx) error {
	// 	c.JSON()
	// 	return c.SendString("Hello World!")
	// })

	log.Fatal(app.Listen(cfg.HTTPServer.Host + ":" + cfg.HTTPServer.Port))
}
