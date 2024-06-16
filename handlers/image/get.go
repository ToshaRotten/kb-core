package image

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"main/models/database"
	"main/models/response"
)

type getImageResponse struct {
	Items []database.Image `json:"items"`
	response.Response
}

func GetImage(c fiber.Ctx, db *gorm.DB) error {
	var items []database.Image

	err := db.Model(&database.Image{}).Find(&items).Error
	if err != nil {
		return c.JSON(getImageResponse{nil, response.Error()})
	}

	return c.JSON(getImageResponse{items, response.OK()})
}
