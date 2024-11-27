package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()

	logFile, err := os.OpenFile("api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		Logger.Fatal(err)
	}

	Logger.SetOutput(logFile)
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetLevel(logrus.InfoLevel)
}

func LogEvent(level logrus.Level, message string, fields logrus.Fields) {
	Logger.WithFields(fields).Log(level, message)
}
