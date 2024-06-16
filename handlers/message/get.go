package message

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"main/models/database"
	"main/models/response"
)

type getMessageResponse struct {
	Items []database.Message `json:"items"`
	response.Response
}

func GetMessage(c fiber.Ctx, db *gorm.DB) error {
	var items []database.Message

	err := db.Model(&database.Message{}).Find(&items).Error
	if err != nil {
		return c.JSON(getMessageResponse{nil, response.Error()})
	}

	return c.JSON(getMessageResponse{items, response.OK()})
}
