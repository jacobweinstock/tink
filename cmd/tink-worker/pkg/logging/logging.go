package logging

import (
	"github.com/go-logr/logr"
)

// Logger is the interface for logging.
type Logger interface {
	Init() logr.Logger
}
