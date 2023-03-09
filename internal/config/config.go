package config

type Config struct {
	Username string
	Password string
}

func NewConfig() *Config {
	return &Config{}
}
