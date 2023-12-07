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
		host: "localhost",
		port: "5432",
		user: "root",
		pass: "root",
		db:   "postgre_api",
	}
}
