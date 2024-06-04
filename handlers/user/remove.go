package user

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type removeUserRequest struct {
	Lesson database.Lesson `json:"lesson"`
}

type removeUserResponse struct {
	response.Response
}

func RemoveUser(c fiber.Ctx, db *gorm.DB) error {
	var request removeUserRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(removeUserResponse{response.Error()})
	}
	db.Delete(&request.Lesson)

	return c.JSON(removeUserResponse{response.OK()})
}
