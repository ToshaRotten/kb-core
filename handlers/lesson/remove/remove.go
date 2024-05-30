package remove

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type RemoveLessonRequest struct {
	Lesson database.Lesson `json:"lesson"`
}

type RemoveLessonResponse struct {
	response.Response
}

func RemoveLesson(c fiber.Ctx, db *gorm.DB) error {
	var request RemoveLessonRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(RemoveLessonResponse{response.Error()})
	}
	db.Delete(&request.Lesson)

	return c.JSON(RemoveLessonResponse{response.OK()})
}
