package database

type Lesson struct {
	ID                  uint   `gorm:"primaryKey" json:"id"`
	Teacher             User   `gorm:"embedded"   json:"teacher"`
	Group               Group  `gorm:"embedded"   json:"group"`
	Name                string `                  json:"name"`
	HoursTheory         uint   `                  json:"hours_theory"`
	HoursPractice       uint   `                  json:"hours_practice"`
	HoursConsult        uint   `                  json:"hours_consult"`
	HoursExam           uint   `                  json:"hours_exam"`
	FirstSemesterHours  uint   `                  json:"first_semester_hours"`
	SecondSemesterHours uint   `                  json:"second_semester_hours"`
	ThirdSemesterHours  uint   `                  json:"third_semester_hours"`
	ForthSemesterHours  uint   `                  json:"forth_semester_hours"`
	FifthSemesterHours  uint   `                  json:"fifth_semster_hours"`
	SixthSemesterHours  uint   `                  json:"sixtx_semster_hours"`
	HoursTotal          uint   `                  json:"hours_total"`
	CountSubgroups      uint   `                  json:"count_subgroups"`
	Faculty             uint   `                  json:"faculty"`
}
