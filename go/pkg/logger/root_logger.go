package logger

import (
	"context"
	"sync/atomic"
)

var rootLogger atomic.Pointer[Logger] //nolint:gochecknoglobals

func init() { //nolint:gochecknoinits
	rootLogger.Store(New())
}

// SetRootLogger ルートロガーを設定します
func SetRootLogger(logger *Logger) {
	rootLogger.Store(logger)
}

// Debug 開発時にデバッグ情報を出力する
func Debug(ctx context.Context, msg string, fields ...Field) {
	rootLogger.Load().Debug(ctx, msg, fields...)
}

// Info システムが正常な動作や主要なイベント毎に出力する
func Info(ctx context.Context, msg string, fields ...Field) {
	rootLogger.Load().Info(ctx, msg, fields...)
}

// Warn ユーザの進行に影響がなく、システム的に問題がないがアプリケーションとしては想定外の場合に出力する
func Warn(ctx context.Context, msg string, fields ...Field) {
	rootLogger.Load().Warn(ctx, msg, fields...)
}

// Error ユーザの進行に影響を及ぼすような想定外のエラーを出力します
func Error(ctx context.Context, msg string, fields ...Field) {
	rootLogger.Load().Error(ctx, msg, fields...)
}

// Fatal システム全体への影響を及ぼすエラーを出力します
func Fatal(ctx context.Context, msg string, fields ...Field) {
	rootLogger.Load().Fatal(ctx, msg, fields...)
}
