package logger

import (
	"github.com/sirupsen/logrus"
	"log"
)

var logger *logrus.Logger

func InitLogger(logLevel string) {
	if logger != nil {
		log.Fatalln("attempt to reinitialize logger")
	}
	logger = &logrus.Logger{}
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatalln(err)
	}
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetOutput(log.Writer())
}

func Get() *logrus.Logger {
	if logger == nil {
		log.Fatalln("Logger is not initialized")
	}
	return logger
}
