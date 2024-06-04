package lesson

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type updateLessonRequest struct {
	Lesson database.Lesson `json:"lesson"`
}

type updateLessonResponse struct {
	response.Response
}

func UpdateLesson(c fiber.Ctx, db *gorm.DB) error {
	var request updateLessonRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(updateLessonResponse{response.Error()})
	}
	db.Save(&request.Lesson)

	return c.JSON(updateLessonResponse{response.OK()})
}
