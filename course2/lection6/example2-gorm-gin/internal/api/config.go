package api

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	bindAddr string `toml:"web_server.bind_addr"`
	dbUri    string `toml:"storage.database_uri"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Init() error {
	_, err := toml.DecodeFile("./config/config.toml", c)
	return err
}
