package remove

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type RemoveUserRequest struct {
	Lesson database.Lesson `json:"lesson"`
}

type RemoveUserResponse struct {
	response.Response
}

func RemoveUser(c fiber.Ctx, db *gorm.DB) error {
	var request RemoveUserRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(RemoveUserResponse{response.Error()})
	}
	db.Delete(&request.Lesson)

	return c.JSON(RemoveUserResponse{response.OK()})
}
