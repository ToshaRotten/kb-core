package delete

import (
	"main/models/database"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type DeleteGroupRequest struct {
	database.Group
}

type DeleteGroupResponse struct{}

func DeleteGroup(c fiber.Ctx, db *gorm.DB) error {
	err := c.JSON(c.Body())
	if err != nil {
		return fiber.ErrBadRequest
	}
	return nil
}
