package database

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	Level       string `json:"level"`
	Description string `json:"description"`
	SenderRefer int
	Sender      Service `json:"sender_service" gorm:"foreignKey:SenderRefer"`
}
