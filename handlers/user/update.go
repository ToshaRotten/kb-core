package user

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type updateUserRequest struct {
	Lesson database.Lesson `json:"lesson"`
}

type updateUserResponse struct {
	response.Response
}

func UpdateUser(c fiber.Ctx, db *gorm.DB) error {
	var request updateUserRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(updateUserResponse{response.Error()})
	}
	db.Save(&request.Lesson)

	return c.JSON(updateUserResponse{response.OK()})
}
