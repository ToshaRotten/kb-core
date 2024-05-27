package database

type Group struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `                  json:"code"`
	Professor Professor `gorm:"embedded"   json:"professor"`
	Price     uint      `                  json:"price"`
}
