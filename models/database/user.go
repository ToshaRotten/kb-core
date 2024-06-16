package database

import "gorm.io/gorm"

// User
type User struct {
	gorm.Model
	Name  string  `                  json:"name"`
	Email *string `                  json:"email"`
	Role  uint    `                  json:"role"`
}
