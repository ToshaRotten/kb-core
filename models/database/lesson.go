package database

type Lesson struct {
	ID            uint  `gorm:"primaryKey" json:"id"`
	Teacher       User  `gorm:"embedded"   json:"teacher"`
	Group         Group `gorm:"embedded"   json:"group"`
	HoursTheory   uint  `                  json:"hours_theory"`
	HoursPractice uint  `                  json:"hours_practice"`
}
