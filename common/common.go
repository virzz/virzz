package common

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// Logger 日志工具
	Logger *logrus.Logger

	// DebugMode -
	DebugMode bool = false
)

func init() {
	debugEnv := os.Getenv("DEBUG")
	if debugEnv != "" && debugEnv != "0" && debugEnv != "false" {
		DebugMode = true
	}
	// Debug Mode
	level := logrus.InfoLevel
	if DebugMode {
		level = logrus.DebugLevel
	}
	Logger = InitLogger("virzz.log", level)
}
