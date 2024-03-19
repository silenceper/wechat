package util

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger = logger{}

type logger struct {
	debug bool
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// SetDebug sets the debug flag
func (r logger) SetDebug(isDebug bool) {
	r.debug = isDebug
}

// Debugf logs a message at level Debug on the standard logger.
func (r logger) Debugf(format string, args ...interface{}) {
	if r.debug {
		log.Debugf(format, args...)
	}
}

// Error logs a message at level Error on the standard logger.
func (r logger) Error(args ...interface{}) {
	if r.debug {
		log.Error(args...)
	}
}
