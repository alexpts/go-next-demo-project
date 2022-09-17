package logger

import "context"

type ILogger interface {
	Debug(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Warning(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	//WithField(ctx context.Context, k string, v interface{}) context.Context
	//WithFields(ctx context.Context, fields map[string]interface{}) context.Context
	//WithError(ctx context.Context, err error) context.Context
	//WithTag(ctx context.Context, k, v string) context.Context
}
