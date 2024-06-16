package service

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"main/models/database"
	"main/models/response"
)

type getServiceResponse struct {
	Items []database.Service `json:"items"`
	response.Response
}

func GetService(c fiber.Ctx, db *gorm.DB) error {
	var items []database.Service

	err := db.Model(&database.Service{}).Find(&items).Error
	if err != nil {
		return c.JSON(getServiceResponse{nil, response.Error()})
	}

	return c.JSON(getServiceResponse{items, response.OK()})
}
