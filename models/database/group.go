package database

type Group struct {
	ID        uint `gorm:"primaryKey"`
	Code      string
	Professor Professor `gorm:"embedded"`
	Price     uint
}
