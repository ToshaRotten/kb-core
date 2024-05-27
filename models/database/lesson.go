package database

type Lesson struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Professor Professor `gorm:"embedded"   json:"professor"`
	Group     Group     `gorm:"embedded"   json:"group"`
}
