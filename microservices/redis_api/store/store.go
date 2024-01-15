package store

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Store struct {
	config       *Config
	db           *redis.Client
	StudentStore *StudentStore
}

func New() *Store {
	return &Store{
		config: NewConfig(),
	}
}

func (s *Store) Open() error {
	db := redis.NewClient(&redis.Options{
		Addr:     s.config.addr,
		Password: s.config.pass,
		DB:       s.config.db,
	})

	if err := db.Ping(ctx).Err(); err != nil {
		fmt.Println("connect error")
		return err
	}

	s.db = db

	fmt.Println("redis connection opened")
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
