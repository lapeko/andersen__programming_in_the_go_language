package api

import (
	"github.com/gorilla/mux"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/storage"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *API) Start() error {
	if err := a.configureLogger(); err != nil {
		return err
	}
	a.configureRouting()
	if err := a.configureStorage(); err != nil {
		return err
	}
	a.logger.Debug("DB successfully connected")
	a.logger.Debug("Server is running on port: ", a.config.BindAddr)
	return http.ListenAndServe(":"+strconv.Itoa(a.config.BindAddr), a.router)
}
