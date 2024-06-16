package image

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"main/models/database"
	"main/models/response"
)

type createImageRequest struct {
	Image database.Image `json:"image"`
}

type createImageResponse struct {
	response.Response
}

func CreateImage(c fiber.Ctx, db *gorm.DB) error {
	var request createImageRequest
	if err := json.Unmarshal(c.Body(), &request.Image); err != nil {
		return c.JSON(createImageResponse{response.Error()})
	}

	db.Create(&request.Image)
	return c.JSON(createImageResponse{response.OK()})
}
