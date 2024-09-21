package api

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	WebServer struct {
		BindAddr int `toml:"bind_addr"`
	} `toml:"web_server"`
	Storage struct {
		DbUri string `toml:"db_uri"`
	} `toml:"storage"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Init() (err error) {
	_, err = toml.DecodeFile("./config/config.toml", c)
	return
}
