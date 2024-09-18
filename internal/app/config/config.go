package config

type Config struct {
	Port uint
}

func New() *Config {
	return &Config{
		Port: 8080,
	}
}
