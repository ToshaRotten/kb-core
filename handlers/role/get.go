package role

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"main/models/database"
	"main/models/response"
)

type getRoleResponse struct {
	Items []database.Role `json:"items"`
	response.Response
}

func GetRole(c fiber.Ctx, db *gorm.DB) error {
	var items []database.Role

	err := db.Model(&database.Role{}).Find(&items).Error
	if err != nil {
		return c.JSON(getRoleResponse{nil, response.Error()})
	}

	return c.JSON(getRoleResponse{items, response.OK()})
}
