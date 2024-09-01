package api

import (
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	BindAddr int    `toml:"bind_addr" env:"BIND_ADDR"`
	LogLevel string `toml:"log_level" env:"LOG_LEVEL"`
}

func NewConfig(typeOfFile string, path string) *Config {
	config := Config{
		BindAddr: 8080,
		LogLevel: "Debug",
	}
	switch typeOfFile {
	case "toml":
		if path == "" {
			path = "./configs/api.toml"
		}
		if _, err := toml.DecodeFile(path, &config); err != nil {
			log.Fatalln(err)
		}
	case "env":
		if path == "" {
			path = "./configs/.env"
		}
		if err := godotenv.Load(path); err != nil {
			log.Fatalf("No .env file found: %v", err)
		}
		if err := env.Parse(&config); err != nil {
			log.Fatalf("Failed to parse environment variables: %v", err)
		}
	default:
		log.Fatalln(errors.New("unknown type of app config"))
	}
	return &config
}
