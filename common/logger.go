package common

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

// LogHook to send error logs via bot.
type LogHook struct {
	Field  string
	Skip   int
	levels []logrus.Level
}

func findCaller(skip int) string {
	file := ""
	line := 0
	for i := 0; i < 10; i++ {
		file, line = getCaller(skip + i)
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func getCaller(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0
	}
	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line
}

// Fire - LogHook::Fire
func (hook LogHook) Fire(entry *logrus.Entry) error {
	if entry.Level == logrus.ErrorLevel || entry.Level == logrus.DebugLevel {
		entry.Data[hook.Field] = findCaller(hook.Skip)
	}
	return nil
}

// Levels - LogHook::Levels
func (hook LogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// NewLogHook - Create a hook to be added to an instance of logger
func newLogHook(levels ...logrus.Level) logrus.Hook {
	hook := LogHook{
		Field:  "line",
		Skip:   5,
		levels: levels,
	}
	if len(hook.levels) == 0 {
		hook.levels = logrus.AllLevels
	}
	return &hook
}

// InitLogger -
func InitLogger(filename string, level logrus.Level) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: false,
		TimestampFormat:        "2006-01-02 15:04:05",
		FullTimestamp:          true,
	})
	logger.SetOutput(os.Stdout)
	logger.AddHook(newLogHook())
	return logger
}
