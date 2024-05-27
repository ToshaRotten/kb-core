package database

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `                  json:"first_name"`
	LastName  string `                  json:"last_name"`
	ThirdName string `                  json:"third_name"`
}
