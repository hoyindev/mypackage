package jsonlog

import (
	"fmt"
	"io"
	"runtime"

	"github.com/sirupsen/logrus"
)

//JSONLogger object with logrus.Entry and config for line number
type JSONLogger struct {
	DefaultLogger  *logrus.Entry
	withLineNumber bool
}

//NewJSONLogger create JSONLogger object
func NewJSONLogger() *JSONLogger {
	logger := JSONLogger{
		DefaultLogger:  logrus.WithFields(logrus.Fields{}),
		withLineNumber: false,
	}
	return &logger
}

//SetJSONFormatter set log in json format
func (logger *JSONLogger) SetJSONFormatter() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

//SetOutput set io writer
func (logger *JSONLogger) SetOutput(out io.Writer) {
	logrus.SetOutput(out)
}

//AddFields with more fields
func (logger *JSONLogger) AddFields(fields logrus.Fields) {
	logger.DefaultLogger = logger.DefaultLogger.WithFields(fields)
}

//SetLineNumber set line number in log
func (logger *JSONLogger) SetLineNumber(with bool) {
	logger.withLineNumber = with
}

func (logger *JSONLogger) lineNumberMiddleware() *logrus.Entry {
	if logger.withLineNumber {
		return logger.DefaultLogger.WithFields(
			logrus.Fields{
				"caller": getLineNumber(3),
			},
		)
	}
	return logger.DefaultLogger

}

func getLineNumber(deep int) string {
	pc, fn, line, _ := runtime.Caller(deep)
	return fmt.Sprintf("[Caller] [%s:%s:%d]", runtime.FuncForPC(pc).Name(), fn, line)
}

//SetLevel set level
func (logger *JSONLogger) SetLevel(level logrus.Level) {
	logger.DefaultLogger.Logger.SetLevel(level)
}

//GetLevel get level
func (logger *JSONLogger) GetLevel() logrus.Level {
	return logger.DefaultLogger.Logger.GetLevel()
}

//Info print info
func (logger *JSONLogger) Info(args ...interface{}) {
	logger.lineNumberMiddleware().Info(args...)
}

//Infoln println info
func (logger *JSONLogger) Infoln(args ...interface{}) {
	logger.lineNumberMiddleware().Infoln(args...)
}

//Infof printf info
func (logger *JSONLogger) Infof(format string, args ...interface{}) {
	logger.lineNumberMiddleware().Infof(format, args...)
}

//Warn print warm
func (logger *JSONLogger) Warn(args ...interface{}) {
	logger.lineNumberMiddleware().Warn(args...)
}

//Warnln println warn
func (logger *JSONLogger) Warnln(args ...interface{}) {
	logger.lineNumberMiddleware().Warnln(args...)
}

//Warnf printf warn
func (logger *JSONLogger) Warnf(format string, args ...interface{}) {
	logger.lineNumberMiddleware().Warnf(format, args...)
}

//Trace print trace
func (logger *JSONLogger) Trace(args ...interface{}) {
	logger.lineNumberMiddleware().Trace(args...)
}

//Traceln println trace
func (logger *JSONLogger) Traceln(args ...interface{}) {
	logger.lineNumberMiddleware().Traceln(args...)
}

//Tracef print tracef
func (logger *JSONLogger) Tracef(format string, args ...interface{}) {
	logger.lineNumberMiddleware().Tracef(format, args...)
}

//Debug print debug
func (logger *JSONLogger) Debug(args ...interface{}) {
	logger.lineNumberMiddleware().Debug(args...)
}

//Debugln println debug
func (logger *JSONLogger) Debugln(args ...interface{}) {
	logger.lineNumberMiddleware().Debugln(args...)
}

//Debugf printf debug
func (logger *JSONLogger) Debugf(format string, args ...interface{}) {
	logger.lineNumberMiddleware().Debugf(format, args...)
}

//Error print err
func (logger *JSONLogger) Error(args ...interface{}) {
	logger.lineNumberMiddleware().Error(args...)
}

//Errorln println err
func (logger *JSONLogger) Errorln(args ...interface{}) {
	logger.lineNumberMiddleware().Errorln(args...)
}

//Errorf printf err
func (logger *JSONLogger) Errorf(format string, args ...interface{}) {
	logger.lineNumberMiddleware().Errorf(format, args...)
}

//Fatal will exit status 1
func (logger *JSONLogger) Fatal(args ...interface{}) {
	logger.lineNumberMiddleware().Fatal(args...)
}

//Panic will exit status 1
func (logger *JSONLogger) Panic(args ...interface{}) {
	logger.lineNumberMiddleware().Panic(args...)
}
