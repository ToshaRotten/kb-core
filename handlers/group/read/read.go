package get

import (
	"main/models/database"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetGroupByID(c fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id", "0")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fiber.ErrBadRequest
	}

	group := database.Group{
		ID: uint(idInt),
	}

	db.Take(group)
	return c.JSON(group)
}
