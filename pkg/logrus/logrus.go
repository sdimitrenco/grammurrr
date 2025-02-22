package logrus

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	lgr "github.com/sirupsen/logrus"
)
type LogrusLogger struct {
	entry *lgr.Entry
}

type writerHook struct {
	Writer []io.Writer
	LogLevels []lgr.Level
}

func (hook *writerHook) Fire(entry *lgr.Entry) error {
	line, err := entry.String()

	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}

	return err
}

func (hook *writerHook) Levels() []lgr.Level {
	return hook.LogLevels
}

func NewLogrusLogger() *LogrusLogger {
	l := lgr.New()
	l.SetReportCaller(true)
	l.Formatter = &lgr.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	_ = os.MkdirAll("logs", 0755)
	allFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	l.SetOutput(io.Discard)
	l.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: lgr.AllLevels,
	})

	return &LogrusLogger{entry: lgr.NewEntry(l)}
}

// Реализация методов логирования
func (l *LogrusLogger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l *LogrusLogger) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l *LogrusLogger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}
