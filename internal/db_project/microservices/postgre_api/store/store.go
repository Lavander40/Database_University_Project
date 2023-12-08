package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
	StudentStore *StudentStore
}

func New() *Store {
	return &Store{
		config: NewConfig(),
	}
}

func (s *Store) Open() error {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", s.config.host, s.config.port, s.config.user, s.config.pass, s.config.db)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	s.db = db
	defer db.Close()

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
