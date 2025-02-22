package logging

type LoggerInterface interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
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

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}


