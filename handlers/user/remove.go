package user

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type removeUserRequest struct {
	User database.User `json:"user"`
}

type removeUserResponse struct {
	response.Response
}

func RemoveUser(c fiber.Ctx, db *gorm.DB) error {
	var request removeUserRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(removeUserResponse{response.Error()})
	}
	db.Delete(&request.User)

	return c.JSON(removeUserResponse{response.OK()})
}
