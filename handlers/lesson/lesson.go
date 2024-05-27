package lesson

import (
	"main/models/database"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetLessonByID(c fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id", "0")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fiber.ErrBadRequest
	}

	lesson := database.Lesson{
		ID: uint(idInt),
	}

	db.Take(lesson)
	return c.JSON(lesson)
}
