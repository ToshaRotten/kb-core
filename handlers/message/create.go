package message

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"main/models/database"
	"main/models/response"
)

type createMessageRequest struct {
	Message database.Message `json:"message"`
}

type createMessageResponse struct {
	response.Response
}

func CreateMessage(c fiber.Ctx, db *gorm.DB) error {
	var request createMessageRequest

	if err := json.Unmarshal(c.Body(), &request.Message); err != nil {
		return c.JSON(createMessageResponse{response.Error()})
	}
	db.Create(&request.Message)

	return c.JSON(createMessageResponse{response.OK()})
}
