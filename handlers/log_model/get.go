package log_model

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"main/models/database"
	"main/models/response"
)

type getLogResponse struct {
	Items []database.Log `json:"items"`
	response.Response
}

func GetLog(c fiber.Ctx, db *gorm.DB) error {
	var items []database.Log

	err := db.Model(&database.Log{}).Find(&items).Error
	if err != nil {
		return c.JSON(getLogResponse{nil, response.Error()})
	}

	return c.JSON(getLogResponse{items, response.OK()})
}
