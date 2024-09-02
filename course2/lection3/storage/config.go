package storage

type Config struct {
	DBUri string `toml:"database_uri" env:"DATABASE_URI"`
}

func NewConfig() *Config {
	return &Config{}
}
