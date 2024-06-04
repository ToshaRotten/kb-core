package lesson

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type removeLessonRequest struct {
	Lesson database.Lesson `json:"lesson"`
}

type removeLessonResponse struct {
	response.Response
}

func RemoveLesson(c fiber.Ctx, db *gorm.DB) error {
	var request removeLessonRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(removeLessonResponse{response.Error()})
	}
	db.Delete(&request.Lesson)

	return c.JSON(removeLessonResponse{response.OK()})
}
