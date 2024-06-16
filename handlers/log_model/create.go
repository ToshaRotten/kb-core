package log_model

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"main/models/database"
	"main/models/response"
)

type createLogRequest struct {
	Log database.Log `json:"log"`
}

type createLogResponse struct {
	response.Response
}

func CreateLog(c fiber.Ctx, db *gorm.DB) error {
	var request createLogRequest
	if err := json.Unmarshal(c.Body(), &request.Log); err != nil {
		return c.JSON(createLogResponse{response.Error()})
	}
	db.Create(&request.Log)

	return c.JSON(createLogResponse{response.OK()})
}
