package group

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type updateGroupRequest struct {
	database.Group `json:"group"`
}

type updateGroupResponse struct {
	response.Response
}

func UpdateGroup(c fiber.Ctx, db *gorm.DB) error {
	var request updateGroupRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(updateGroupResponse{response.Error()})
	}
	db.Save(&request.Group)

	return c.JSON(updateGroupResponse{response.OK()})
}
