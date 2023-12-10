package store

import "db_project/internal/db_project/microservices/postgre_api/model"

type LessonStore struct {
	store *Store
}

func (s *LessonStore) Get(lessonId int) (model.Lesson, error) {
	var lesson model.Lesson

	row := s.store.db.QueryRow("SELECT * FROM lessons WHERE id = $1", lessonId)
	if err := row.Scan(&lesson.Id, &lesson.Type, &lesson.Date, &lesson.EquipId, &lesson.CourseId); err != nil {
		return lesson, err
	}

	return lesson, nil
}

func (s *LessonStore) GetAll() ([]model.Lesson, error) {
	var lessons []model.Lesson

	rows, err := s.store.db.Query("SELECT * FROM lessons")
	if err != nil {
		return lessons, err
	}
	defer rows.Close()

	for rows.Next() {
		lesson := model.Lesson{}
		err := rows.Scan(&lesson.Id, &lesson.Type, &lesson.Date, &lesson.EquipId, &lesson.CourseId)
		if err != nil {
			return lessons, err
		}
		lessons = append(lessons, lesson)
	}

	return lessons, nil
}