package store

type Config struct {
	addr string
	pass string
	db   int
	// stud_db int
	// equip_db int
	// aud_db int
}

func NewConfig() *Config {
	return &Config{
		addr: "localhost:6379",
		pass: "",
		db:   0,
		// stud_db:   0,
		// equip_db:   1,
		// aud_db:   2,
	}
}
