package model

type Attendance struct {
	Id       string `json:"id"`
	StudentId   string `json:"stud_id"`
	ScheduleId string `json:"sched_id"`
}

type Lesson struct {
	Id      string `json:"id"`
	Type    string `json:"typing"`
	Date    string `json:"date"`
	LectionID string `json:"lection_id"`
	EquipId string `json:"equip_id"`
	CourseId string `json:"course_id"`
}

type Student struct {
	Id      string `json:"id"`
	Name    string `json:"full_name"`
	GroupId string `json:"group_id"`
}

type Schedule struct {
	Id       string `json:"id"`
	GroupId  string `json:"group_id"`
	LessonId string `json:"lesson_id"`
	AudId    string `json:"room_id"`
}

type Equipment struct {
	Id       string `json:"id"`
	Name string `json:"equip"`
}

type Group struct {
	Id       string `json:"id"`
	DepartId string `json:"depart_id"`
	SpecId   string `json:"spec_id"`
}

type Room struct {
	Id      string `json:"id"`
	Number  string `json:"number"`
}

type Course struct {
	Id       string `json:"id"`
	Name string `json:"naming"`
	Description   string `json:"description"`
}

type Speciality struct {
	Id       string `json:"id"`
	Name string `json:"naming"`
	Code   string `json:"code"`
	DepartId string `json:"depart_id"`
}

type Speciality_Course struct {
	Id       string `json:"id"`
	SpecId   string `json:"spec_id"`
	CourseIdId string `json:"course_id"`
}

type Rate struct{
	Id int
	Score float64
}
