package lesson

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type createLessonRequest struct {
	database.Lesson `json:"lesson"`
}

type createLessonResponse struct {
	response.Response
}

func CreateLesson(c fiber.Ctx, db *gorm.DB) error {
	var request createLessonRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(createLessonResponse{response.Error()})
	}
	db.Create(&request.Group)

	return c.JSON(createLessonResponse{response.OK()})
}
