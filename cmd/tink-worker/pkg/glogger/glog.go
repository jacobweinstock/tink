package glogger

import (
	"github.com/go-logr/glogr"
	"github.com/go-logr/logr"
)

type Logger struct {
	Glog logr.Logger
}

// Init returns a new  glogr.
func Init() *Logger {
	return &Logger{Glog: glogr.New()}
}
