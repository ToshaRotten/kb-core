package create

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type CreateLessonRequest struct {
	database.Lesson `json:"lesson"`
}

type CreateLessonResponse struct {
	response.Response
}

func CreateLesson(c fiber.Ctx, db *gorm.DB) error {
	var request CreateLessonRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(CreateLessonResponse{response.Error()})
	}
	db.Create(&request.Group)

	return c.JSON(CreateLessonResponse{response.OK()})
}
