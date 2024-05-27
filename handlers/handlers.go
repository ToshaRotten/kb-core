package handlers

import (
	"main/models/database"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetUserByID(c fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id", "0")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fiber.ErrBadRequest
	}

	user := database.User{
		ID: uint(idInt),
	}

	db.Take(&user)

	return c.JSON(user)
}
