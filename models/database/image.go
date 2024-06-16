package database

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Name    string `                  json:"name"`
	Version string `                  json:"version"`
}
