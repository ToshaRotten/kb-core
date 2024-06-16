package user

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"main/models/database"
	"main/models/response"
)

type updateUserRequest struct {
	User database.User `json:"user"`
}

type updateUserResponse struct {
	response.Response
}

func UpdateUser(c fiber.Ctx, db *gorm.DB) error {
	var request updateUserRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(updateUserResponse{response.Error()})
	}
	db.Save(request.User)

	return c.JSON(updateUserResponse{response.OK()})
}
