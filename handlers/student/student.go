package student

import (
	"main/models/database"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetStudentByID(c fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id", "0")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fiber.ErrBadRequest
	}

	student := database.Student{
		ID: uint(idInt),
	}

	db.Take(student)
	return c.JSON(student)
}
