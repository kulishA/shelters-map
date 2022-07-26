package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Logger
}

func NewLogger() *Logger {
	log := logrus.New()

	return &Logger{log: log}
}

func (l *Logger) Warn(message string) {
	l.log.Warning(message)
}

func (l *Logger) Info(message string) {
	l.log.Info(message)
}

func (l *Logger) Error(message string, err error) {
	l.log.Error(fmt.Printf("%s - %s", message, err.Error()))
}
