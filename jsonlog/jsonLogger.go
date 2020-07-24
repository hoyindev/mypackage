package jsonlog

import (
	"os"

	"github.com/sirupsen/logrus"
)

type JsonLogger struct {
	DefaultLogger *logrus.Entry
}

func NewJsonLogger() *JsonLogger {
	requestLogger := logrus.WithFields(logrus.Fields{})
	logger := JsonLogger{
		DefaultLogger: requestLogger,
	}
	return &logger
}

func (logger *JsonLogger) SetUpConfig() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	// logrus.SetLevel(log.WarnLevel)
}

func (logger *JsonLogger) SetDefaultMsg(fields logrus.Fields) {
	logger.DefaultLogger = logger.DefaultLogger.WithFields(fields)
}

func (logger *JsonLogger) SetLevel(level logrus.Level) {
	logger.DefaultLogger.Logger.SetLevel(level)
}

func (logger *JsonLogger) GetLevel() logrus.Level {
	return logger.DefaultLogger.Logger.GetLevel()
}

func (logger *JsonLogger) Info(args ...interface{}) {
	logger.DefaultLogger.Info(args...)
}

func (logger *JsonLogger) Infoln(args ...interface{}) {
	logger.DefaultLogger.Infoln(args...)
}

func (logger *JsonLogger) Infof(format string, args ...interface{}) {
	logger.DefaultLogger.Infof(format, args...)
}

func (logger *JsonLogger) Warn(args ...interface{}) {
	logger.DefaultLogger.Warn(args...)
}

func (logger *JsonLogger) Warnln(args ...interface{}) {
	logger.DefaultLogger.Warnln(args...)
}

func (logger *JsonLogger) Warnf(format string, args ...interface{}) {
	logger.DefaultLogger.Warnf(format, args...)
}

func (logger *JsonLogger) Trace(args ...interface{}) {
	logger.DefaultLogger.Trace(args...)
}

func (logger *JsonLogger) Traceln(args ...interface{}) {
	logger.DefaultLogger.Traceln(args...)
}

func (logger *JsonLogger) Tracef(format string, args ...interface{}) {
	logger.DefaultLogger.Tracef(format, args...)
}

func (logger *JsonLogger) Debug(args ...interface{}) {
	logger.DefaultLogger.Debug(args...)
}

func (logger *JsonLogger) Debugln(args ...interface{}) {
	logger.DefaultLogger.Debugln(args...)
}

func (logger *JsonLogger) Debugf(format string, args ...interface{}) {
	logger.DefaultLogger.Debugf(format, args...)
}

func (logger *JsonLogger) Error(args ...interface{}) {
	logger.DefaultLogger.Error(args...)
}

func (logger *JsonLogger) Errorln(args ...interface{}) {
	logger.DefaultLogger.Errorln(args...)
}

func (logger *JsonLogger) Errorf(format string, args ...interface{}) {
	logger.DefaultLogger.Errorf(format, args...)
}

//will exit status 1
func (logger *JsonLogger) Fatal(args ...interface{}) {
	logger.DefaultLogger.Fatal(args...)
}

//will exit status 1
func (logger *JsonLogger) Panic(args ...interface{}) {
	logger.DefaultLogger.Panic(args...)
}

//below are examples on how to create log func with extra param (if extra fields required to be added)
func (logger *JsonLogger) CustomInfo(user string, args ...interface{}) {
	logger.DefaultLogger.WithFields(
		logrus.Fields{
			"user": user,
		},
	).Info(args...)
}

func (logger *JsonLogger) CustomWarn(user string, args ...interface{}) {
	logger.DefaultLogger.WithFields(
		logrus.Fields{
			"user": user,
		},
	).Warn(args...)
}
