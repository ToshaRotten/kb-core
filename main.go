package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"main/config"
	"main/handlers/image"
	"main/handlers/log_model"
	"main/handlers/message"
	"main/handlers/role"
	"main/handlers/service"
	"main/handlers/user"
	"main/models/database"
	"main/models/response"
)

func main() {
	logrus.Info("starting app ...")
	cfg := config.MustLoadConfig("configs/local.env")
	app := fiber.New()
	app.Use(cors.New())

	connectionString := "postgresql://" + cfg.Storage.Username + ":" + cfg.Storage.Password + "@" + cfg.Storage.Host + ":" + cfg.Storage.Port + "/" + cfg.Storage.Database
	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		logrus.Error(err.Error())
	}

	logrus.Info("auto migrate ...")
	err = db.AutoMigrate(
		&database.Image{},
		&database.Log{},
		&database.Role{},
		&database.Service{},
		&database.Message{},
		&database.User{},
	)
	if err != nil {
		logrus.Error(err.Error())
	}

	app.Use(func(c fiber.Ctx) error {
		logrus.Info(
			" " + string(
				c.Method(),
			) + string(
				c.Request().RequestURI(),
			) + " " + c.Request().
				String(),
		)
		return c.Next()
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

	app.Get("/health/", func(c fiber.Ctx) error {
		return c.JSON(response.OK)
	})

	app.Post("/log/", func(c fiber.Ctx) error {
		return log_model.CreateLog(c, db)
	})

	app.Get("/log/", func(c fiber.Ctx) error {
		return log_model.GetLog(c, db)
	})

	app.Post("/service/", func(c fiber.Ctx) error {
		return service.CreateService(c, db)
	})

	app.Get("/service/", func(c fiber.Ctx) error {
		return service.GetService(c, db)
	})

	app.Post("/message/", func(c fiber.Ctx) error {
		return message.CreateMessage(c, db)
	})

	app.Get("/message/", func(c fiber.Ctx) error {
		return message.GetMessage(c, db)
	})

	app.Post("/role/", func(c fiber.Ctx) error {
		return role.CreateRole(c, db)
	})

	app.Get("/role/", func(c fiber.Ctx) error {
		return role.GetRole(c, db)
	})

	app.Post("/image/", func(c fiber.Ctx) error {
		return image.CreateImage(c, db)
	})

	app.Get("/image/", func(c fiber.Ctx) error {
		return image.GetImage(c, db)
	})

	logrus.Info("starting server ...")
	log.Fatal(app.Listen(cfg.HTTPServer.Host + ":" + cfg.HTTPServer.Port))
}
