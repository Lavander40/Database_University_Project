package store

import "db_project/internal/db_project/microservices/postgre_api/model"

type CourseStore struct {
	store *Store
}

func (s *CourseStore) Get(courseId int) (model.Course, error) {
	var course model.Course

	row := s.store.db.QueryRow("SELECT * FROM courses WHERE id = $1", courseId)
	if err := row.Scan(&course.Id, &course.Name, &course.Description); err != nil {
		return course, err
	}

	return course, nil
}

func (s *CourseStore) GetAll() ([]model.Course, error) {
	var courses []model.Course

	rows, err := s.store.db.Query("SELECT * FROM courses")
	if err != nil {
		return courses, err
	}
	defer rows.Close()

	for rows.Next() {
		course := model.Course{}
		err := rows.Scan(&course.Id, &course.Name, &course.Description)
		if err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}