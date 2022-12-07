package logger

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gookit/color"
	"github.com/virzz/virzz/common"
)

// const (
// 	PREFIX_SUCCESS = "+"
// 	PREFIX_ERROR   = "-"
// 	PREFIX_WARN    = "!"
// 	PREFIX_INFO    = "~"
// 	PREFIX_DEBUG   = "*"
// )

var std = DefaultOutput()

func DefaultOutput() *log.Logger {
	return log.New(os.Stderr, "", 0)
}

func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

func SetPrefix(prefix ...string) {
	if len(prefix) > 0 && prefix[0] != "" {
		std.SetPrefix(color.LightMagenta.Sprintf("[%s] ", prefix[0]))
	} else {
		std.SetPrefix("")
	}
}

func Panic(v ...any) {
	std.Panic(v...)
}

func Print(v ...any) {
	std.Output(3, fmt.Sprint(v...))
}

func Printf(format string, v ...any) {
	Print(fmt.Sprintf(format, v...))
}

func Success(v ...any) {
	Print(color.LightGreen.Sprint(append([]any{"[+] "}, v...)...))
}
func SuccessF(format string, v ...any) {
	Print(color.LightGreen.Sprintf("[+] "+format, v...))
}

func Error(v ...any) {
	Print(color.LightRed.Sprint(append([]any{"[-] "}, v...)...))
}
func ErrorF(format string, v ...any) {
	Print(color.LightRed.Sprintf("[-] "+format, v...))
}

func Warn(v ...any) {
	Print(color.LightYellow.Sprint(append([]any{"[!] "}, v...)...))
}
func WarnF(format string, v ...any) {
	Print(color.LightYellow.Sprintf("[!] "+format, v...))
}

func Info(v ...any) {
	Print(color.LightCyan.Sprint(append([]any{"[~] "}, v...)...))
}
func InfoF(format string, v ...any) {
	Print(color.LightCyan.Sprintf("[~] "+format, v...))
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
		Print(color.LightBlue.Sprint(append([]any{"\n[*] "}, v...)...))
		std.SetFlags(0)
	}
}
func DebugF(format string, v ...any) {
	if common.DebugMode {
		std.SetFlags(log.Lshortfile | log.LstdFlags)
		Print(color.LightBlue.Sprintf("\n[*] "+format, v...))
		std.SetFlags(0)
	}
}

// ======= Log to File =========

type LogWriter struct {
	lf *os.File
}

func (lw LogWriter) Write(p []byte) (n int, err error) {
	if lw.lf != nil {
		lw.lf.Write(p)
	}
	return os.Stderr.Write(p)
}

func FileOutput(name string) *log.Logger {
	lf, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		Error(err)
	}
	var lw io.Writer = LogWriter{lf: lf}
	return log.New(lw, "", 0)
}

func LogAsFileOutput(name ...string) {
	_name := "./virzz.log"
	if len(name) > 0 && len(name[0]) > 0 {
		_name = name[0]
	}
	std = FileOutput(_name)
}
