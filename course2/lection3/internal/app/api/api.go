package api

import (
	"github.com/gorilla/mux"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/pkg/logger"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/storage"
	"net/http"
	"strconv"
)

type API struct {
	config  *Config
	router  *mux.Router
	storage *storage.Storage
}

func New(config *Config) *API {
	return &API{
		config: config,
		router: mux.NewRouter(),
	}
}

func (a *API) Start() error {
	log := logger.Get()
	if err := a.configureStorage(); err != nil {
		return err
	}
	a.configureRouting()
	log.Debug("DB successfully connected")
	log.Debug("Server is running on port: ", a.config.BindAddr)
	return http.ListenAndServe(":"+strconv.Itoa(a.config.BindAddr), a.router)
}
