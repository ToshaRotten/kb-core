package user

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"main/models/database"
	"main/models/response"
)

type getUserResponse struct {
	Items []database.User `json:"items"`
	response.Response
}

func GetUser(c fiber.Ctx, db *gorm.DB) error {
	var items []database.User

	err := db.Model(&database.User{}).Find(&items).Error
	if err != nil {
		return c.JSON(getUserResponse{nil, response.Error()})
	}

	return c.JSON(getUserResponse{items, response.OK()})
}
