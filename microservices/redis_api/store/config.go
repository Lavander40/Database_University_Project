package store

type Config struct {
	addr string
	pass string
	db   int
}

func NewConfig() *Config {
	return &Config{
		addr: "r_db:6379",
		pass: "",
		db:   0,
	}
}
