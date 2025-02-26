package logging

type LoggerInterface interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Debug(args ...interface{})
	WithField(key string, value interface{}) *Logger
}

type Logger struct {
	logger LoggerInterface
}

func NewLogger(l LoggerInterface) *Logger {
	return &Logger{logger: l}
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	newLogger := l.logger.WithField(key, value)
	l.logger = newLogger.logger
	return newLogger
}
