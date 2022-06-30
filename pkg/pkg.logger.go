package pkg

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

var logger = &Logger{}

func SetupLogger() {
	logger = &Logger{logrus.New()}
	logger.Formatter = &logrus.JSONFormatter{}
	logrus.SetOutput(os.Stdout)

}

func GetLogger() *Logger {
	return logger
}
