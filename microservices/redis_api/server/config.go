package server

type Config struct {
	BindAddr string
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":4041",
	}
}
