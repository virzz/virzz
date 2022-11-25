// package logger
package logger

import (
	"testing"
)

func TestDebug(t *testing.T) {
	Debug("test")
	DebugF("test Debug")
}

func TestSuccess(t *testing.T) {
	Success("test")
	SuccessF("test Success")
}

func TestError(t *testing.T) {
	Error("test")
	ErrorF("test Error")
}

func TestWarn(t *testing.T) {
	Warn("test")
	WarnF("test Warn")
}

func TestInfo(t *testing.T) {
	Info("test")
	InfoF("test Info")
}

func TestPrint(t *testing.T) {
	Print("test")
	Printf("test Print")
}

func TestPrefix(t *testing.T) {
	SetPrefix("PREFIX")
	Print("test")
	Printf("test Prefix")
	Debug(`Debug
ttttttbikasebd
aegvagsrrdgvs
	agrehted`)
	SetPrefix()
}

func TestPanic(t *testing.T) {
	Panic("test")
}
