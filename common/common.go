package common

import (
	"os"

	"github.com/virzz/logger"
)

const (
	Author = "陌竹"
	Email  = "mozhu233@outlook.com"
)

var (
	DebugMode bool   = false
	Mode      string = "dev"
)

func init() {
	debugEnv := os.Getenv("VIRZZ_DEBUG")
	if debugEnv == "true" || debugEnv == "1" || debugEnv == "on" {
		DebugMode = true
		logger.SetDebug(true)
	}
	if Mode == "dev" {
		DebugMode = true
	}
	// Force off debug mode
	if debugEnv == "off" {
		DebugMode = false
	}
}
