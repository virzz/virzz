package common

import (
	"os"
)

const (
	Author = "陌竹(@mozhu1024)"
	Email  = "mozhu233@outlook.com"
)

var (
	DebugMode bool = false
)

func init() {
	debugEnv := os.Getenv("DEBUG")
	if debugEnv == "true" || debugEnv == "1" || debugEnv == "on" {
		DebugMode = true
	}
	// Force off debug mode
	if debugEnv == "off" {
		DebugMode = false
	}
}
