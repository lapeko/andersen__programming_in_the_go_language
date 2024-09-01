package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type API struct {
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
	}
}

func (a *API) Start() error {
	if err := a.configureLogger(); err != nil {
		return err
	}
	fmt.Println("Server is running")
	return nil
}
