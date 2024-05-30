package database

type Group struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	Code             string `                  json:"code"`
	Name             string `                  json:"name"`
	ClassroomTeacher User   `gorm:"embedded"   json:"classroom_teacher"`
}
