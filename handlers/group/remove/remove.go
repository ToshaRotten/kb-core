package remove

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type RemoveGroupRequest struct {
	database.Group `json:"group"`
}

type RemoveGroupResponse struct {
	response.Response
}

func RemoveGroup(c fiber.Ctx, db *gorm.DB) error {
	var request RemoveGroupRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(RemoveGroupResponse{response.Error()})
	}
	db.Delete(&request.Group)

	return c.JSON(RemoveGroupResponse{response.OK()})
}
