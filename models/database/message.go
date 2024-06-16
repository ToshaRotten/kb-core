package database

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Type         string `json:"type"`
	Description  string `json:"description"`
	SenderRefer  int
	ReciverRefer int
	Sender       Service `json:"sender_service"  gorm:"foreignKey:SenderRefer"`
	Reciver      Service `json:"reciver_service" gorm:"foreignKey:ReciverRefer"`
}
