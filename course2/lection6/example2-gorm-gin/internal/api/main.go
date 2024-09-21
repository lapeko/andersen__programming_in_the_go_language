package api

import (
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection6/example2-gorm-gin/storage"
	"log"
)

type Api struct {
	config  *Config
	Storage *storage.Storage
}

func New() *Api {
	return &Api{}
}

func (a *Api) Init() {
	a.config = NewConfig()
	if err := a.config.Init(); err != nil {
		log.Fatalf("Config initialization error: %v\n", err)
	}
	a.Storage = storage.New(a.config.dbUri)
	if err := a.Storage.Init(); err != nil {
		log.Fatalf("Storage initialization error: %v\n", err)
	}
	if err := a.Storage.Migrate(); err != nil {
		log.Fatalf("Storage migration error: %v\n", err)
	}
}
