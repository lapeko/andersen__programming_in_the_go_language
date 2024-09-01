package api

type Config struct {
	BindAddr int `toml:"BIND_ADDR"`
}

func NewConfig() *Config {
	return &Config{}
}
