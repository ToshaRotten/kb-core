package create

import (
	"encoding/json"
	"main/models/database"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type CreateGroupRequest struct {
	database.Group `json:"group"`
}

type CreateGroupResponse struct{}

func CreateGroup(c fiber.Ctx, db *gorm.DB) error {
	var request CreateGroupRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return err
	}
	return &fiber.InvalidUnmarshalError{}
}
