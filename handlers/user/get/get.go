package get

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type GetUserRequest struct {
	User database.User `json:"user"`
}

type GetUserResponse struct {
	Items []database.User `json:"items"`
	response.Response
}

func GetGroups(c fiber.Ctx, db *gorm.DB) error {
	var request GetUserRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(GetUserResponse{nil, response.Error()})
	}
	var items []database.User

	err := db.Model(&database.Group{}).Find(&items).Error
	if err != nil {
		return c.JSON(GetUserResponse{nil, response.Error()})
	}

	return c.JSON(GetUserResponse{items, response.OK()})
}

func GetUserByID(c fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id", "0")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fiber.ErrBadRequest
	}

	user := database.User{
		ID: uint(idInt),
	}

	db.Take(user)
	return c.JSON(user)
}
