package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    *string
	Age      uint8
	Birthday *time.Time
	Role     uint
}
