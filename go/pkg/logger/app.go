package logger

import (
	"io"
	"log/slog"
	"time"
)

type AppHandler struct {
	*slog.JSONHandler
}

// NewAppHandler アプリケーションログ用のハンドラー（AppHandler）を生成します
func NewAppHandler(writer io.Writer) *AppHandler {
	replaceAttr := func(_ []string, attr slog.Attr) slog.Attr {
		switch attr.Key {
		case slog.LevelKey:
			return extendLogLevelAttr(attr, slog.LevelKey)
		case slog.MessageKey:
			return slog.Attr{Key: "message", Value: attr.Value}
		}
		return attr
	}
	handlerOptions := &slog.HandlerOptions{
		AddSource:   false,
		Level:       slogLeveler(),
		ReplaceAttr: replaceAttr,
	}
	return &AppHandler{
		slog.NewJSONHandler(writer, handlerOptions),
	}
}

// NewAppHandlerWithTime 出力時刻が与えられた time に固定される AppHandler を生成します
func NewAppHandlerWithTime(writer io.Writer, time time.Time) *AppHandler {
	replaceAttr := func(_ []string, attr slog.Attr) slog.Attr {
		switch attr.Key {
		case slog.LevelKey:
			return extendLogLevelAttr(attr, slog.LevelKey)
		case slog.MessageKey:
			return slog.Attr{Key: "message", Value: attr.Value}
		case slog.TimeKey:
			return slog.Attr{Key: slog.TimeKey, Value: slog.TimeValue(time)}
		}
		return attr
	}
	handlerOptions := &slog.HandlerOptions{
		AddSource:   false,
		Level:       slogLeveler(),
		ReplaceAttr: replaceAttr,
	}
	return &AppHandler{
		slog.NewJSONHandler(writer, handlerOptions),
	}
}
