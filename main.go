package main

import (
	"log"
	"log/slog"
	"main/config"
	"main/handlers/group"
	"main/handlers/lesson"
	"main/handlers/user"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	slog.Info("starting app ...")
	cfg := config.MustLoadConfig("configs/local.env")
	app := fiber.New()
	app.Use(cors.New())

	connectionString := "postgresql://" + cfg.Storage.Username + ":" + cfg.Storage.Password + "@" + cfg.Storage.Host + ":" + cfg.Storage.Port + "/" + cfg.Storage.Database
	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info("auto migrate ...")
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

	// Group
	app.Post("/group/", func(c fiber.Ctx) error {
		return group.CreateGroup(c, db)
	})

	app.Delete("/group/", func(c fiber.Ctx) error {
		return group.RemoveGroup(c, db)
	})

	app.Put("/group/", func(c fiber.Ctx) error {
		return group.UpdateGroup(c, db)
	})

	app.Get("/group/", func(c fiber.Ctx) error {
		return group.GetGroups(c, db)
	})

	app.Get("/group/:id", func(c fiber.Ctx) error {
		return group.GetGroupByID(c, db)
	})

	// Lesson
	app.Get("/lesson/", func(c fiber.Ctx) error {
		return lesson.GetLessons(c, db)
	})

	app.Post("/lesson/", func(c fiber.Ctx) error {
		return lesson.CreateLesson(c, db)
	})

	app.Delete("/lesson/", func(c fiber.Ctx) error {
		return lesson.RemoveLesson(c, db)
	})

	app.Put("/lesson/", func(c fiber.Ctx) error {
		return lesson.UpdateLesson(c, db)
	})

	app.Get("/lesson/:id", func(c fiber.Ctx) error {
		return lesson.GetLessonByID(c, db)
	})

	// User
	app.Get("/user/", func(c fiber.Ctx) error {
		return user.GetUser(c, db)
	})

	app.Post("/user/", func(c fiber.Ctx) error {
		return user.CreateUser(c, db)
	})

	app.Delete("/user/", func(c fiber.Ctx) error {
		return user.RemoveUser(c, db)
	})

	app.Put("/user/", func(c fiber.Ctx) error {
		return user.UpdateUser(c, db)
	})

	app.Get("/user/:id", func(c fiber.Ctx) error {
		return user.GetUserByID(c, db)
	})

	app.Get("/health/", func(c fiber.Ctx) error {
		return c.JSON(response.OK)
	})

	slog.Info("starting server ...")
	log.Fatal(app.Listen(cfg.HTTPServer.Host + ":" + cfg.HTTPServer.Port))
}
