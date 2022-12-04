// package logger
package logger

import (
	"fmt"
	"testing"

	"github.com/gookit/color"
)

func TestDebug(t *testing.T) {
	Debug("test")
	DebugF("test Debug")
}

func TestSuccess(t *testing.T) {
	Success("test")
	Success("test", "aaa")
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

// func TestPanic(t *testing.T) {
// 	Panic("test")
// }

func BenchmarkSprintf(b *testing.B) {
	b.Run("Sprintf - 1", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			color.LightRed.Sprintf("[-] %s", "test")
		}
	})
	b.Run("Sprintf - 2", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			color.LightGreen.Sprintf(fmt.Sprintf("%s %%s", "+"), "test")
		}
	})
}

func BenchmarkAppend(b *testing.B) {
	v := []any{"a", "b"}
	b.Run("append-1", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			color.LightGreen.Sprint(append([]any{"[+] "}, v...)...)
		}
	})
	b.Run("append-2", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			color.LightRed.Sprintf("[-] %s", v...)
		}
	})
}
