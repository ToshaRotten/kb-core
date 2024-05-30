package update

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type UpdateGroupRequest struct {
	database.Group `json:"group"`
}

type UpdateGroupResponse struct {
	response.Response
}

func UpdateGroup(c fiber.Ctx, db *gorm.DB) error {
	var request UpdateGroupRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(UpdateGroupResponse{response.Error()})
	}
	db.Save(&request.Group)

	return c.JSON(UpdateGroupResponse{response.OK()})
}
