package main

import (
	"log"
	"log/slog"
	"main/config"
	"main/handlers"
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
	//
	// storage, err := storage.New(cfg.Storage)
	// if err != nil {
	// 	slog.Error(err.Error())
	// }

	connectionString := "postgresql://" + cfg.Storage.Username + ":" + cfg.Storage.Password + "@" + cfg.Storage.Host + ":" + cfg.Storage.Port + "/" + cfg.Storage.Database
	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		slog.Error(err.Error())
	}

	// gormDB, err := gorm.Open(postgres.New(postgres.Config{
	// 	Conn: storage.DB,
	// }), &gorm.Config{})
	// if err != nil {
	// 	slog.Error(err.Error())
	// }

	err = db.AutoMigrate(
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

	app.Get("/user/:id", handlers.GetUserByID(fiber.Ctx{}, db))
	log.Fatal(app.Listen(cfg.HTTPServer.Host + ":" + cfg.HTTPServer.Port))
}
