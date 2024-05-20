package database

type Group struct {
	ID        uint `gorm:"primaryKey;default:auto_random()"`
	Code      string
	Professor Professor `gorm:"embedded"`
	Price     uint
}
