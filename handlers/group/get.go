package group

import (
	"main/models/database"
	"main/models/response"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type getGroupResponse struct {
	Items []database.Group `json:"items"`
	response.Response
}

func GetGroups(c fiber.Ctx, db *gorm.DB) error {
	var items []database.Group

	err := db.Model(&database.Group{}).Find(&items).Error
	if err != nil {
		return c.JSON(getGroupResponse{nil, response.Error()})
	}

	return c.JSON(getGroupResponse{items, response.OK()})
}

func GetGroupByID(c fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id", "0")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fiber.ErrBadRequest
	}

	group := database.Group{
		ID: uint(idInt),
	}

	db.Take(group)
	return c.JSON(group)
}
