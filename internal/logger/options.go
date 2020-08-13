package logger

import "github.com/sirupsen/logrus"

type Option func(*Logger)

func WithLevel(level string) Option {
	return func(l *Logger) {
		if parseLevel, err := logrus.ParseLevel(level); err == nil {
			l.l.SetLevel(parseLevel)
			l.l.WithField("log_level", parseLevel).Debugln("Set logger level")
		}
	}
}

func WithJsonFormatter(prettyPrint bool) Option {
	return func(l *Logger) {
		l.Logger().SetFormatter(&logrus.JSONFormatter{PrettyPrint: prettyPrint})
	}
}

func WithCustomFormatter(formatter logrus.Formatter) Option {
	return func(l *Logger) {
		l.Logger().SetFormatter(formatter)
	}
}
