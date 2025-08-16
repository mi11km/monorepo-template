package logger

import (
	"io"
	"log/slog"
)

type ConsoleHandler struct {
	*slog.TextHandler
}

// NewConsoleHandler cli ツールなどのコンソールにログ出力する ConsoleHandler を生成します
func NewConsoleHandler(writer io.Writer) *ConsoleHandler {
	replaceAttr := func(_ []string, attr slog.Attr) slog.Attr {
		if attr.Key == slog.LevelKey {
			return extendLogLevelAttr(attr, slog.LevelKey)
		}
		return attr
	}

	handlerOptions := &slog.HandlerOptions{
		AddSource:   false,
		Level:       slogLeveler(),
		ReplaceAttr: replaceAttr,
	}
	return &ConsoleHandler{
		slog.NewTextHandler(writer, handlerOptions),
	}
}
