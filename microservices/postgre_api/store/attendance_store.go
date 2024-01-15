package store

import (
	"postgre_api/model"
	"github.com/lib/pq"
)

type AttendanceStore struct {
	store *Store
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

func (s *AttendanceStore) GetRate(lessons, students []int) ([]model.Rate, error) {
	var rates []model.Rate
	
	rows, err := s.store.db.Query("SELECT students.id, (SELECT count(id) FROM attendances WHERE stud_id = students.id AND sched_id IN (SELECT id FROM schedules WHERE lesson_id = ANY($1)))/(SELECT count(id) FROM lessons WHERE id = ANY($1))::float AS score FROM students WHERE id = ANY($2) ORDER BY score ASC LIMIT 10;", pq.Array(lessons), pq.Array(students))
	if err != nil {
		return rates, err
	}
	defer rows.Close()

	for rows.Next(){
        rate := model.Rate{}
        err := rows.Scan(&rate.Id, &rate.Score)
        if err != nil{
            return rates, err
        }
        rates = append(rates, rate)
    }
	return rates, nil
}