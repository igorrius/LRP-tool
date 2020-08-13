package logger

import "context"

type contextKeyType int

var contextKey contextKeyType

func NewContext(ctx context.Context, l *Logger) context.Context {
	return context.WithValue(ctx, contextKey, l)
}

func FromContext(ctx context.Context) (*Logger, bool) {
	l, ok := ctx.Value(contextKey).(*Logger)
	return l, ok
}

//FromContextOrNew return logger from context or create a new one
func FromContextOrNew(ctx context.Context) *Logger {
	if l, ok := FromContext(ctx); ok {
		return l
	}
	return NewLogger()
}
