package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
	EquipmentStore *EquipmentStore
	RoomStore *RoomStore
	CourseStore *CourseStore
	SpecialityStore *SpecialityStore
	GroupStore *GroupStore
	StudentStore *StudentStore
	LessonStore *LessonStore
	ScheduleStore *ScheduleStore
	AttendanceStore *AttendanceStore
}

func New() *Store {
	return &Store{
		config: NewConfig(),
	}
}

func (s *Store) Open() error {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", s.config.user, s.config.pass, s.config.host, s.config.port, s.config.db)
	//connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", s.config.host, s.config.port, s.config.user, s.config.pass, s.config.db)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	s.db = db

	fmt.Println("postgre connection opened")
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Student() *StudentStore {
	if s.StudentStore != nil {
		return s.StudentStore
	}

	s.StudentStore = &StudentStore{
		store: s,
	}
	return s.StudentStore
}

func (s *Store) Equip() *EquipmentStore {
	if s.EquipmentStore != nil {
		return s.EquipmentStore
	}

	s.EquipmentStore = &EquipmentStore{
		store: s,
	}
	return s.EquipmentStore
}

func (s *Store) Room() *RoomStore {
	if s.RoomStore != nil {
		return s.RoomStore
	}

	s.RoomStore = &RoomStore{
		store: s,
	}
	return s.RoomStore
}

func (s *Store) Course() *CourseStore {
	if s.CourseStore != nil {
		return s.CourseStore
	}

	s.CourseStore = &CourseStore{
		store: s,
	}
	return s.CourseStore
}

func (s *Store) Spec() *SpecialityStore {
	if s.SpecialityStore != nil {
		return s.SpecialityStore
	}

	s.SpecialityStore = &SpecialityStore{
		store: s,
	}
	return s.SpecialityStore
}

func (s *Store) Group() *GroupStore {
	if s.GroupStore != nil {
		return s.GroupStore
	}

	s.GroupStore = &GroupStore{
		store: s,
	}
	return s.GroupStore
}

func (s *Store) Lesson() *LessonStore {
	if s.LessonStore != nil {
		return s.LessonStore
	}

	s.LessonStore = &LessonStore{
		store: s,
	}
	return s.LessonStore
}

func (s *Store) Schedule() *ScheduleStore {
	if s.ScheduleStore != nil {
		return s.ScheduleStore
	}

	s.ScheduleStore = &ScheduleStore{
		store: s,
	}
	return s.ScheduleStore
}

func (s *Store) Attend() *AttendanceStore {
	if s.AttendanceStore != nil {
		return s.AttendanceStore
	}

	s.AttendanceStore = &AttendanceStore{
		store: s,
	}
	return s.AttendanceStore
}
