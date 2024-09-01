package api

import (
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
