package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log *logrus.Logger

func Init(level logrus.Level){
	log = logrus.New()
	log.Out = os.Stdout
	log.SetLevel(level)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Error(args ...interface{}){
	log.Error(args...)
}