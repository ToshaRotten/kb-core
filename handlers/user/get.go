package user

import (
	"main/models/database"
	"main/models/response"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type getUserResponse struct {
	Items []database.User `json:"items"`
	response.Response
}

func GetUser(c fiber.Ctx, db *gorm.DB) error {
	var items []database.User

	err := db.Model(&database.User{}).Find(&items).Error
	if err != nil {
		return c.JSON(getUserResponse{nil, response.Error()})
	}

	return c.JSON(getUserResponse{items, response.OK()})
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
	return c.JSON(getUserResponse{[]database.User{user}, response.OK()})
}
