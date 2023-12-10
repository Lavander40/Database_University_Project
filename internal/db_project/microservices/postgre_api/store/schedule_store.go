package store

import "db_project/internal/db_project/microservices/postgre_api/model"

type ScheduleStore struct {
	store *Store
}

func (s *ScheduleStore) Get(scheduleId int) (model.Schedule, error) {
	var schedule model.Schedule

	row := s.store.db.QueryRow("SELECT * FROM schedules WHERE id = $1", scheduleId)
	if err := row.Scan(&schedule.Id, &schedule.GroupId, &schedule.LessonId, &schedule.AudId); err != nil {
		return schedule, err
	}

	return schedule, nil
}

func (s *ScheduleStore) GetAll() ([]model.Schedule, error) {
	var schedules []model.Schedule

	rows, err := s.store.db.Query("SELECT * FROM schedules")
	if err != nil {
		return schedules, err
	}
	defer rows.Close()

	for rows.Next() {
		schedule := model.Schedule{}
		err := rows.Scan(&schedule.Id, &schedule.GroupId, &schedule.LessonId, &schedule.AudId)
		if err != nil {
			return schedules, err
		}
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}