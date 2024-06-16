package user

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"main/models/database"
	"main/models/response"
)

type createUserRequest struct {
	User database.User `json:"user"`
}

type createUserResponse struct {
	response.Response
}

func CreateUser(c fiber.Ctx, db *gorm.DB) error {
	var request createUserRequest
	if err := json.Unmarshal(c.Body(), &request.User); err != nil {
		return c.JSON(createUserResponse{response.Error()})
	}
	db.Create(&request.User)

	return c.JSON(createUserResponse{response.OK()})
}
