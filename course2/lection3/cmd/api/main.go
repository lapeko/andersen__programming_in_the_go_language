package main

import (
	"flag"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/internal/app/api"
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
	a := api.New(config)
	if err := a.Start(); err != nil {
		log.Fatalln(err)
	}
}
