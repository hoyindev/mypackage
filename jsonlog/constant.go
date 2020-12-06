package jsonlog

import (
	"github.com/sirupsen/logrus"
)

const (
	// Panic level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	Panic = logrus.PanicLevel
	// Fatal level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	Fatal = logrus.FatalLevel
	// Error level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	Error = logrus.ErrorLevel
	// Warn level. Non-critical entries that deserve eyes.
	Warn = logrus.WarnLevel
	// Info level. General operational entries about what's going on inside the
	// application.
	Info = logrus.InfoLevel
	// Debug level. Usually only enabled when debugging. Very verbose logging.
	Debug = logrus.DebugLevel
	// Trace level. Designates finer-grained informational events than the Debug.
	Trace = logrus.TraceLevel
)
