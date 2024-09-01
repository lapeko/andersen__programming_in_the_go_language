package api

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{}
}
