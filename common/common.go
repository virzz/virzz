package common

import (
	"os"
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
	debugEnv := os.Getenv("MOZHU1024")
	if debugEnv == "dev" || Mode == "dev" {
		DebugMode = true
	}
}
