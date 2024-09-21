package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection6/example2-gorm-gin/internal/routes"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection6/example2-gorm-gin/storage"
	"log"
)

type Api struct {
	config  *Config
	Router  *gin.Engine
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
	a.Storage = storage.New(a.config.Storage.DbUri)
	if err := a.Storage.Init(); err != nil {
		log.Fatalf("Storage initialization error: %v\n", err)
	}
	if err := a.Storage.Migrate(); err != nil {
		log.Fatalf("Storage migration error: %v\n", err)
	}
}

func (a *Api) Start() {
	a.Router = gin.Default()
	routes.SetupArticlesRouter("/articles", a.Router, a.Storage.Repo.Articles)
	if err := a.Router.Run(fmt.Sprintf(":%d", a.config.WebServer.BindAddr)); err != nil {
		log.Fatalf("Starting server error: %v\n", err)
	}
}

func (a *Api) Close() {
	a.Storage.Close()
}
