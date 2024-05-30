package update

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type UpdateLessonRequest struct {
	Lesson database.Lesson `json:"lesson"`
}

type UpdateLessonResponse struct {
	response.Response
}

func UpdateLesson(c fiber.Ctx, db *gorm.DB) error {
	var request UpdateLessonRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(UpdateLessonResponse{response.Error()})
	}
	db.Save(&request.Lesson)

	return c.JSON(UpdateLessonResponse{response.OK()})
}
