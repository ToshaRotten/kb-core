package service

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"main/models/database"
	"main/models/response"
)

type createServiceRequest struct {
	Service database.Service `json:"serivce"`
}

type createServiceResponse struct {
	response.Response
}

func CreateService(c fiber.Ctx, db *gorm.DB) error {
	var request createServiceRequest
	if err := json.Unmarshal(c.Body(), &request.Service); err != nil {
		return c.JSON(createServiceResponse{response.Error()})
	}
	db.Create(&request.Service)

	return c.JSON(createServiceResponse{response.OK()})
}
