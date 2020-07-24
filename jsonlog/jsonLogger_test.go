package jsonlog

import (
	"testing"

	"fmt"
)

func TestJsonLogger(t *testing.T) {
	myLogger := NewJsonLogger()
	fields := make(map[string]interface{})
	fields["service"] = "api"

	myLogger.SetUpConfig()
	myLogger.SetDefaultMsg(fields)
	currentLevel := myLogger.GetLevel()
	fmt.Println("default level", currentLevel, " now set trace level")
	myLogger.SetLevel(TraceLevel)

	myLogger.CustomInfo("userB", "msg in log info")
	myLogger.CustomWarn("userA", "msg in log warn")

	myLogger.Info("info msg,", "Gopher")
	myLogger.Infoln("infoln msg,", "Gopher")
	myLogger.Infof("%s", "normal msg")

	myLogger.Warn("warn msg")
	myLogger.Warnln("warnln msg,", "Gopher")
	myLogger.Warnf("%s", "warnf msg,")

	myLogger.Error("error msg,", "Gopher")
	myLogger.Errorln("errorln msg,", "Gopher")
	myLogger.Errorf("%s", "errorf msg")

	//***********************************************
	//set debug level to show debug level log || set trace level to show debug and trace level log
	//***********************************************
	// myLogger.SetLevel(logrus.TraceLevel)
	// currentLevel = myLogger.GetLevel()
	// fmt.Println("after set level", currentLevel)

	myLogger.Debug("debug msg,", "Gopher")
	myLogger.Debugln("debugln msg,", "Gopher")
	myLogger.Debugf("%s", "tracef msg")

	myLogger.Trace("trace msg,", "Gopher")
	myLogger.Traceln("traceln msg,", "Gopher")
	myLogger.Tracef("%s", "tracef msg")

	t.Error("done")
}
