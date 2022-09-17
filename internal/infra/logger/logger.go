package logger

import (
	"context"
	"fmt"
)

type Logger struct{}

func (l *Logger) Debug(ctx context.Context, args ...interface{}) {
	_ = fmt.Sprint(args...)
}

func (l *Logger) Info(ctx context.Context, args ...interface{}) {
	_ = fmt.Sprint(args...)
}

func (l *Logger) Warning(ctx context.Context, args ...interface{}) {
	_ = fmt.Sprint(args...)
}

func (l *Logger) Error(ctx context.Context, args ...interface{}) {
	_ = fmt.Sprint(args...)
}
