package api

import "github.com/sirupsen/logrus"

func (a *API) configureLogger() (err error) {
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err == nil {
		a.logger.SetLevel(level)
	}
	return
}
