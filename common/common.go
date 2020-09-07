package common

import (
	"os"

	log "github.com/virink/vlogger"
)

var (
	// Logger 日志工具
	Logger *log.Logger

	// DebugMode -
	DebugMode bool = false
)

func init() {
	level := log.LevelError
	debugEnv := os.Getenv("DEBUG")
	if debugEnv != "" && debugEnv != "0" && debugEnv != "false" {
		DebugMode = true
		level = log.LevelDebug
	}
	InitLogger(level)
}
