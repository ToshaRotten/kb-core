package database

type Lesson struct {
	ID        uint      `gorm:"primaryKey"`
	Professor Professor `gorm:"embedded"`
	Group     Group     `gorm:"embedded"`
}
