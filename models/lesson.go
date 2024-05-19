package models

type Lesson struct {
	Id             int
	Code           string
	Name           string
	HoursTheory    int
	HoursPractice  int
	HoursConsult   int
	HoursExam      int
	FirstSemestr   int
	SecondSemester int
	ThirdSemester  int
	ForthSemester  int
	FifthSemester  int
	SixthSemester  int
	HoursTotal     int
	CountSubgroups int
	Faculty        string
}
