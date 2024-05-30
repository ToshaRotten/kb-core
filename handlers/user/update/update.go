package update

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type UpdateUserRequest struct {
	Lesson database.Lesson `json:"lesson"`
}

type UpdateUserResponse struct {
	response.Response
}

func UpdateUser(c fiber.Ctx, db *gorm.DB) error {
	var request UpdateUserRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(UpdateUserResponse{response.Error()})
	}
	db.Save(&request.Lesson)

	return c.JSON(UpdateUserResponse{response.OK()})
}
