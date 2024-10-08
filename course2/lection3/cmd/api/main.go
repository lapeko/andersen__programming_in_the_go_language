package main

import (
	"flag"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/internal/app/api"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/pkg/logger"
	"log"
)

var (
	configFormat string
	configPath   string
)

func init() {
	flag.StringVar(&configFormat, "format", "toml", "Type of the config file [env|toml]. Default is \"toml\"")
	flag.StringVar(&configPath, "path", "", "Path to the app config. Default \"./configs/api.toml\" for type toml and \"./configs/.env\" for type env")
}

func main() {
	flag.Parse()
	config := api.NewConfig(configFormat, configPath)
	logger.InitLogger(config.LogLevel)
	a := api.New(config)
	log.Fatalln(a.Start())
}
