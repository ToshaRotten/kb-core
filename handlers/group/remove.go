package group

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type removeGroupRequest struct {
	database.Group `json:"group"`
}

type removeGroupResponse struct {
	response.Response
}

func RemoveGroup(c fiber.Ctx, db *gorm.DB) error {
	var request removeGroupRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(removeGroupResponse{response.Error()})
	}
	db.Delete(&request.Group)

	return c.JSON(removeGroupResponse{response.OK()})
}
