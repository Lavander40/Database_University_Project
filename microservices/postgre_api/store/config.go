package store

type Config struct {
	host string
	port string
	user string
	pass string
	db   string
}

func NewConfig() *Config {
	return &Config{
		host: "p_db",
		port: "5432",
		user: "postgres",
		pass: "postgres",
		db:   "postgres",
	}
}
