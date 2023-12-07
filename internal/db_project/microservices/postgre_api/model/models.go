package model

type Group struct {
	Id       string `json:"id"`
	DepartId string `json:"depart_id"`
	SpecId   string `json:"spec_id"`
}

type Student struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	GroupId string `json:"group_id"`
}

type Lesson struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	EquipId string `json:"equip_id"`
	Date    string `json:"date"`
	GradeId string `json:"grade_id"`
}

type Attendance struct {
	Id      string `json:"id"`
	StudId  string `json:"stud_id"`
	SchedId string `json:"sched_id"`
	Date    string `json:"date"`
}

type Schedule struct {
	Id       string `json:"id"`
	GroupId  string `json:"group_id"`
	LessonId string `json:"lesson_id"`
	AudId    string `json:"aud_id"`
}
