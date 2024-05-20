package database

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	ThirdName string
}
