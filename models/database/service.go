package database

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name        string `json:"name"`
	Status      string `json:"status"`
	Description string `json:"description"`
	AuthorRefer int
	Author      User `json:"author"      gorm:"foreignKey:AuthorRefer"`
}
