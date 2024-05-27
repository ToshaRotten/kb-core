package professor

import (
	"main/models/database"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetProfessorByID(c fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id", "0")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fiber.ErrBadRequest
	}

	professor := database.Professor{
		ID: uint(idInt),
	}

	db.Take(&professor)
	return c.JSON(professor)
}
