package store

import (
	"postgre_api/model"

	"github.com/lib/pq"
)

type AttendanceStore struct {
	store *Store
}

type Rate struct{
	id int
	score int
}

func (s *AttendanceStore) Get(attendanceId int) (model.Attendance, error) {
	var attendance model.Attendance

	row := s.store.db.QueryRow("SELECT * FROM attendances WHERE id = $1", attendanceId)
	if err := row.Scan(&attendance.Id, &attendance.StudentId, &attendance.ScheduleId); err != nil {
		return attendance, err
	}

	return attendance, nil
}

func (s *AttendanceStore) GetAll() ([]model.Attendance, error) {
	var attendances []model.Attendance

	rows, err := s.store.db.Query("SELECT * FROM attendances")
	if err != nil {
		return attendances, err
	}
	defer rows.Close()

	for rows.Next() {
		attendance := model.Attendance{}
		err := rows.Scan(&attendance.Id, &attendance.StudentId, &attendance.ScheduleId)
		if err != nil {
			return attendances, err
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}

func (s *AttendanceStore) GetRate(lessons, students []int) ([]Rate, error) {
	var rates []Rate

	rows, err := s.store.db.Query("select students.id, (Select count(id) From attendances Where stud_id = students.id and sched_id in (SELECT id FROM schedules Where lesson_id in $1))/(Select count(id) from lessons where id In $1)::float as score from students where id in $2 order by score asc limit 10;", pq.Array(lessons), pq.Array(students))
	if err != nil {
		return rates, err
	}
	defer rows.Close()

	for rows.Next(){
        rate := Rate{}
        err := rows.Scan(&rate.id, &rate.score)
        if err != nil{
            return rates, err
        }
        rates = append(rates, rate)
    }

	return rates, nil
}