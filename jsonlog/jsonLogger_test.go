package jsonlog

import (
	"os"
	"testing"

	"fmt"
)

func TestJsonLogger(t *testing.T) {

	myLogger := NewJSONLogger()

	fields := map[string]interface{}{
		"service": "api 1",
		"admin":   "admin A",
	}
	// setup log config
	myLogger.AddFields(fields)
	myLogger.SetLineNumber(true)
	myLogger.SetJSONFormatter()
	myLogger.SetOutput(os.Stdout)

	currentLevel := myLogger.GetLevel()
	fmt.Println("default level", currentLevel, " now set trace level")
	myLogger.SetLevel(Trace)

	myLogger.Info("info msg,", "Gopher")
	myLogger.Infoln("infoln msg,", "Gopher")
	myLogger.Infof("%s", "normal msg")

	myLogger.Warn("warn msg")
	myLogger.Warnln("warnln msg,", "Gopher")
	myLogger.Warnf("%s", "warnf msg,")

	myLogger.Error("error msg,", "Gopher")
	myLogger.Errorln("errorln msg,", "Gopher")
	myLogger.Errorf("%s", "errorf msg")

	myLogger.Debug("debug msg,", "Gopher")
	myLogger.Debugln("debugln msg,", "Gopher")
	myLogger.Debugf("%s", "tracef msg")

	myLogger.Trace("trace msg,", "Gopher")
	myLogger.Traceln("traceln msg,", "Gopher")
	myLogger.Tracef("%s", "tracef msg")

	t.Error("done")
}
