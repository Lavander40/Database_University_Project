package store

import (
	"db_project/internal/db_project/microservices/redis_api/model"
	"encoding/json"
)

type StudentStore struct {
	store *Store
}

func (s *StudentStore) Set(student model.Student) error {
	marshalStudent, err := json.Marshal(student)
	if err != nil {
		return err
	}

	if err := s.store.db.Set(ctx, student.Id, marshalStudent, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (s *StudentStore) Get(studentId string) (string, error) {
	marshalStudent, err := s.store.db.Get(ctx, studentId).Result()
	if err != nil {
		return err.Error(), err
	}

	return marshalStudent, nil
}

// func (s *StudentStore) GetAll() (string, error) {
// 	var studentTemplate model.Student
// 	studentList := []model.Student{}

// 	for _, studentId := range strings.Split(students, ",") {
// 		student := s.store.db.Get(ctx, studentId)

// 		err := student.Err()

// 		if err != nil {
// 			return err.Error(), err
// 		}

// 		studentResult, err := student.Bytes()

// 		if err != nil {
// 			return err.Error(), err
// 		}

// 		err = json.Unmarshal(studentResult, &studentTemplate)

// 		if err != nil {
// 			return err.Error(), err
// 		}

// 		studentList = append(studentList, studentTemplate)
// 	}

// 	marshalledStudentList, err := json.Marshal(studentList)

// 	if err != nil {
// 		return err.Error(), err
// 	}

// 	return string(marshalledStudentList), nil
// }
