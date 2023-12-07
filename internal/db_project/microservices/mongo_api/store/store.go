package store

type Store struct {
	config *Config
}

func New() *Store {
	return &Store{
		config: NewConfig(),
	}
}

func (s *Store) Open() error {
	return nil
}

func (s *Store) Close() {
	
}