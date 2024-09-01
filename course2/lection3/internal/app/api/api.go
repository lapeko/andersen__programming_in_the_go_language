package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type API struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
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
	a.logger.Debug("Server is running on port: ", a.config.BindAddr)
	return http.ListenAndServe(":"+strconv.Itoa(a.config.BindAddr), a.router)
}
