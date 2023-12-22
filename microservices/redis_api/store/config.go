package store

type Config struct {
	addr string
	pass string
	db   int
}

func NewConfig() *Config {
	return &Config{
		addr: "localhost:6379",
		pass: "",
		db:   0,
	}
}
