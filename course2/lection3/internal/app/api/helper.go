package api

import (
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (a *API) configureLogger() (err error) {
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err == nil {
		a.logger.SetLevel(level)
	}
	return
}

func (a *API) configureRouting() {
	a.router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello world"))
	})
}

func (a *API) configureStorage() error {
	a.storage = storage.New(a.config.StorageConfig)
	return a.storage.Open()
}
