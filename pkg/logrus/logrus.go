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
	Writer    []io.Writer
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
	// Убираем SetReportCaller(true), чтобы Logrus не добавлял свои func и file
	l.Formatter = &lgr.TextFormatter{
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

func (l *LogrusLogger) WithCaller() *lgr.Entry {
	return l.entry.WithField("caller", getCaller(4))
}

func getCaller(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}
	fn := runtime.FuncForPC(pc)
	return fmt.Sprintf("%s:%d %s()", path.Base(file), line, path.Base(fn.Name()))
}

func (l *LogrusLogger) Info(args ...interface{}) {
	l.WithCaller().Info(args...)
}

func (l *LogrusLogger) Warn(args ...interface{}) {
	l.WithCaller().Warn(args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	l.WithCaller().Error(args...)
}

func (l *LogrusLogger) Debug(args ...interface{}) {
	l.WithCaller().Debug(args...)
}