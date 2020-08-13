package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	l *logrus.Logger
}

func NewLogger(opts ...Option) *Logger {
	l := &Logger{l: logrus.New()}
	l.applyOptions(opts...)

	return l
}

func (l *Logger) Logger() *logrus.Logger {
	return l.l
}

func (l *Logger) StdLogger() logrus.StdLogger {
	return l.l
}

func (l *Logger) FieldLogger() logrus.FieldLogger {
	return l.l
}

func (l *Logger) Ext1FieldLogger() logrus.Ext1FieldLogger {
	return l.l
}

func (l *Logger) applyOptions(opts ...Option) {
	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated
		// *Logger as the argument
		opt(l)
	}
}
