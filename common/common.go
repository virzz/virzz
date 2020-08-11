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
	debugEnv := os.Getenv("DEBUG")
	if debugEnv != "" && debugEnv != "0" && debugEnv != "false" {
		DebugMode = true
	}
}
