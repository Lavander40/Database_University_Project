package store

import (
	"db_project/internal/db_project/microservices/postgre_api/model"
	"strconv"
)

type StudentStore struct {
	store *Store
}

func (s *StudentStore) Set(student model.Student) error {
	g_id, err := strconv.Atoi(student.GroupId)
    if err != nil {
        return err
    }

	_, err = s.store.db.Exec("INSERT INTO students (full_name, group_id) VALUES ($1, $2)", student.Name, g_id)
	if err != nil {
		return err
	}

	return nil
}

func (s *StudentStore) Get(studentId int) (model.Student, error) {
	var student model.Student

	row := s.store.db.QueryRow("SELECT * FROM students WHERE id = $1", studentId)
	if err := row.Scan(&student.Id, &student.Name, &student.GroupId); err != nil {
		return student, err
	}

	return student, nil
}

func (s *StudentStore) GetAll() ([]model.Student, error) {
	var students []model.Student

	rows, err := s.store.db.Query("SELECT * FROM students")
	if err != nil {
		return students, err
	}
	defer rows.Close()

	for rows.Next(){
        student := model.Student{}
        err := rows.Scan(&student.Id, &student.Name, &student.GroupId)
        if err != nil{
            return students, err
        }
        students = append(students, student)
    }

	return students, nil
}
