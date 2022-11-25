package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/gookit/color"
	"github.com/mozhu1024/virzz/common"
)

var std = log.New(os.Stderr, "", 0)

func SetPrefix(prefix ...string) {
	if len(prefix) > 0 && prefix[0] != "" {
		std.SetPrefix(color.LightMagenta.Sprintf("[%s] ", prefix[0]))
	} else {
		std.SetPrefix("")
	}
}

func Print(v ...any) {
	std.Output(3, fmt.Sprint(v...))
}

func Panic(v ...any) {
	std.Panic(v...)
}

func Printf(format string, v ...any) {
	Print(fmt.Sprintf(format, v...))
}

func Success(v ...any) {
	Print(color.LightGreen.Sprint(v...))
}
func SuccessF(format string, v ...any) {
	Printf(color.LightGreen.Sprintf(format, v...))
}

func Error(v ...any) {
	Print(color.LightRed.Sprint(v...))
}
func ErrorF(format string, v ...any) {
	Print(color.LightRed.Sprintf(format, v...))
}

func Warn(v ...any) {
	Print(color.LightYellow.Sprint(v...))
}
func WarnF(format string, v ...any) {
	Print(color.LightYellow.Sprintf(format, v...))
}

func Info(v ...any) {
	Print(color.LightCyan.Sprintf("[Info] %s", v...))
}
func InfoF(format string, v ...any) {
	Print(color.LightCyan.Sprintf("[Info] "+format, v...))
}

func Normal(v ...any) {
	Print(color.LightWhite.Sprint(v...))
}
func NormalF(format string, v ...any) {
	Print(color.LightWhite.Sprintf(format, v...))
}

func Debug(v ...any) {
	if common.DebugMode {
		std.SetFlags(log.Lshortfile | log.LstdFlags)
		Print("\n", color.LightBlue.Sprint(v...))
		std.SetFlags(0)
	}
}
func DebugF(format string, v ...any) {
	if common.DebugMode {
		std.SetFlags(log.Lshortfile | log.LstdFlags)
		Print("\n", color.LightBlue.Sprintf(format, v...))
		std.SetFlags(0)
	}
}
