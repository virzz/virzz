package common

import (
	log "github.com/virink/vlogger"
)

// InitLogger -
func InitLogger(level int) *log.Logger {
	l := log.NewLogger().SetLevel(level).SetCallDepth(2)
	Logger = l
	return l
}
