package role

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"main/models/database"
	"main/models/response"
)

type createRoleRequest struct {
	Role database.Role `json:"role"`
}

type createRoleResponse struct {
	response.Response
}

func CreateRole(c fiber.Ctx, db *gorm.DB) error {
	var request createRoleRequest

	if err := json.Unmarshal(c.Body(), &request.Role); err != nil {
		return c.JSON(createRoleResponse{response.Error()})
	}
	db.Create(&request.Role)

	return c.JSON(createRoleResponse{response.OK()})
}
