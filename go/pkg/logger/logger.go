package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"
)

const (
	// LogLevelEnv [debug, info, warn, error, fatal] に対応, デフォルトは info
	LogLevelEnv = "LOG_LEVEL"

	// LogFormatEnv [console, json] に対応, デフォルトは json
	LogFormatEnv = "LOG_FORMAT"
)

type Handler = slog.Handler

type Logger struct {
	logger *slog.Logger
}

// New os.Stderr にログを出力する Logger を生成します
func New() *Logger {
	return NewWithWriter(os.Stderr)
}

// NewWithWriter 指定した io.Writer にログを出力する Logger を生成します
func NewWithWriter(writer io.Writer) *Logger {
	var handler slog.Handler
	switch strings.ToLower(os.Getenv(LogFormatEnv)) {
	case "console":
		handler = NewConsoleHandler(writer)
	case "json":
		handler = NewAppHandler(writer)
	default:
		handler = NewAppHandler(writer)
	}

	return &Logger{logger: slog.New(handler)}
}

// NewWithHandler 指定した Handler でログを出力する Logger を生成します
func NewWithHandler(handler Handler) *Logger {
	return &Logger{logger: slog.New(handler)}
}

func (l *Logger) Debug(ctx context.Context, msg string, fields ...Field) {
	l.logger.LogAttrs(ctx, slog.LevelDebug, msg, fields...)
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...Field) {
	l.logger.LogAttrs(ctx, slog.LevelInfo, msg, fields...)
}

func (l *Logger) Warn(ctx context.Context, msg string, fields ...Field) {
	l.logger.LogAttrs(ctx, slog.LevelWarn, msg, fields...)
}

func (l *Logger) Error(ctx context.Context, msg string, fields ...Field) {
	l.logger.LogAttrs(ctx, slog.LevelError, msg, fields...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, fields ...Field) {
	l.logger.LogAttrs(ctx, LevelFatal, msg, fields...)
	os.Exit(1)
}

// Field is an additional logging field.
type Field = slog.Attr

func FieldString(key, val string) Field {
	return slog.String(key, val)
}

func FieldInt(key string, val int) Field {
	return slog.Int(key, val)
}

func FieldInt64(key string, val int64) Field {
	return slog.Int64(key, val)
}

func FieldUint64(key string, val uint64) Field {
	return slog.Uint64(key, val)
}

func FieldFloat64(key string, val float64) Field {
	return slog.Float64(key, val)
}

func FieldBool(key string, val bool) Field {
	return slog.Bool(key, val)
}

func FieldTime(key string, val time.Time) Field {
	return slog.Time(key, val)
}

func FieldDuration(key string, val time.Duration) Field {
	return slog.Duration(key, val)
}

func FieldAny(key string, val any) Field {
	return slog.Any(key, val)
}

func FieldError(err error) Field {
	return slog.String("error", err.Error())
}
