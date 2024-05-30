package create

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type CreateUserRequest struct {
	User database.User `json:"user"`
}

type CreateUserResponse struct {
	response.Response
}

func CreateUser(c fiber.Ctx, db *gorm.DB) error {
	var request CreateUserRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(CreateUserResponse{response.Error()})
	}
	db.Create(&request.User)

	return c.JSON(CreateUserResponse{response.OK()})
}
