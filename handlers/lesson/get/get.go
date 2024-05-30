package get

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type GetLessonRequest struct {
	database.Lesson `json:"lesson"`
}

type GetLessonResponse struct {
	Items []database.Lesson `json:"items"`
	response.Response
}

func GetLessons(c fiber.Ctx, db *gorm.DB) error {
	var request GetLessonRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(GetLessonResponse{nil, response.Error()})
	}
	var items []database.Lesson

	err := db.Model(&database.User{}).Find(&items).Error
	if err != nil {
		return c.JSON(GetLessonResponse{nil, response.Error()})
	}

	return c.JSON(GetLessonResponse{items, response.OK()})
}

func GetLessonByID(c fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id", "0")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fiber.ErrBadRequest
	}

	lesson := database.Lesson{
		ID: uint(idInt),
	}

	db.Take(lesson)
	return c.JSON(lesson)
}
