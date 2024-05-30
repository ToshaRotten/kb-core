package main

import (
	"log"
	"log/slog"
	"main/config"
	"main/handlers/group/create"
	"main/handlers/group/remove"
	"main/handlers/group/update"
	"main/models/database"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.MustLoadConfig("configs/local.env")
	app := fiber.New()
	app.Use(cors.New())

	connectionString := "postgresql://" + cfg.Storage.Username + ":" + cfg.Storage.Password + "@" + cfg.Storage.Host + ":" + cfg.Storage.Port + "/" + cfg.Storage.Database
	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		slog.Error(err.Error())
	}

	err = db.AutoMigrate(
		&database.User{},
		&database.Group{},
		&database.Lesson{},
	)
	if err != nil {
		slog.Error(err.Error())
	}

	app.Use(func(c fiber.Ctx) error {
		slog.Info(
			" " + string(
				c.Method(),
			) + string(
				c.Request().RequestURI(),
			) + " " + c.Request().
				String(),
		)
		return c.Next()
	})

	app.Post("/group/", func(c fiber.Ctx) error {
		return create.CreateGroup(c, db)
	})

	app.Delete("/group/", func(c fiber.Ctx) error {
		return remove.RemoveGroup(c, db)
	})

	app.Put("/group/", func(c fiber.Ctx) error {
		return update.UpdateGroup(c, db)
	})

	log.Fatal(app.Listen(cfg.HTTPServer.Host + ":" + cfg.HTTPServer.Port))
}
