package logrus

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	lgr "github.com/sirupsen/logrus"
	"github.com/sdimitrenco/grammurrr/internal/infrastructure/logging"
)

// LogrusLogger implements the logging.LoggerInterface
type LogrusLogger struct {
	logger *lgr.Logger
	entry  *lgr.Entry
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
		_, err := w.Write([]byte(line))
		if err != nil {
			return err
		}
	}
	return nil
}

func (hook *writerHook) Levels() []lgr.Level {
	return hook.LogLevels
}

// NewLogrusLogger creates a new LogrusLogger instance.
func NewLogrusLogger() *LogrusLogger {
	l := lgr.New()
	l.SetReportCaller(false)
	l.Formatter = &lgr.TextFormatter{
		DisableColors:            false,
		FullTimestamp:            true,
		TimestampFormat:          "2006-01-02T15:04:05-07:00",
		ForceQuote:               true,
		DisableLevelTruncation:   true,
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

	return &LogrusLogger{logger: l, entry: lgr.NewEntry(l)}
}

// WithCaller adds caller information to the log entry.
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

// Info logs an info message.
func (l *LogrusLogger) Info(args ...interface{}) {
	l.entry.WithField("caller", getCaller(4)).Info(args...)
}

// Warn logs a warning message.
func (l *LogrusLogger) Warn(args ...interface{}) {
	l.entry.WithField("caller", getCaller(4)).Warn(args...)
}

// Error logs an error message.
func (l *LogrusLogger) Error(args ...interface{}) {
	l.entry.WithField("caller", getCaller(4)).Error(args...)
}

// Debug logs a debug message.
func (l *LogrusLogger) Debug(args ...interface{}) {
	l.entry.WithField("caller", getCaller(4)).Debug(args...)
}

// Fatal logs a fatal message and exits.
func (l *LogrusLogger) Fatal(args ...interface{}) {
	l.entry.WithField("caller", getCaller(4)).Error(args...)
	l.logger.Exit(1)
}

// WithField adds a field to the log entry.
func (l *LogrusLogger) WithField(key string, value interface{}) *logging.Logger {
	// Create a copy of the current entry and add the new field.
	newEntry := l.entry.WithField(key, value)
	//Create new Logger from new Entry
	newLogger := &LogrusLogger{
		logger: l.logger,
		entry:  newEntry,
	}
	//Return wraped Logger
	return logging.NewLogger(newLogger)
}
