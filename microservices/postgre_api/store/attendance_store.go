package store

import "postgre_api/model"

type AttendanceStore struct {
	store *Store
}

func (s *AttendanceStore) Get(attendanceId int) (model.Attendance, error) {
	var attendance model.Attendance

	row := s.store.db.QueryRow("SELECT * FROM attendances WHERE id = $1", attendanceId)
	if err := row.Scan(&attendance.Id, &attendance.Date, &attendance.StudentId, &attendance.ScheduleId); err != nil {
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
		err := rows.Scan(&attendance.Id, &attendance.Date, &attendance.StudentId, &attendance.ScheduleId)
		if err != nil {
			return attendances, err
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}