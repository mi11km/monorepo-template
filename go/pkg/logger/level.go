package logger

import (
	"log/slog"
	"os"
	"strings"
)

// LevelFatal fatal ログレベル. OpenTelemetry の SeverityNumber を -9 した値 (slog に準拠)
// https://opentelemetry.io/docs/specs/otel/logs/data-model/#field-severitynumber
const LevelFatal slog.Level = 12

func slogLeveler() slog.Leveler {
	leveler := &slog.LevelVar{}
	logLevelEnv := os.Getenv(LogLevelEnv)
	switch strings.ToLower(logLevelEnv) {
	case "fatal":
		// NOTE: ログレベル FATAL 対応
		leveler.Set(LevelFatal)
	default:
		if err := leveler.UnmarshalText([]byte(logLevelEnv)); err != nil {
			// NOTE: デフォルトは info レベルにする
			leveler.Set(slog.LevelInfo)
		}
	}
	return leveler
}

// extendLogLevelAttr ログレベル FATAL 対応用の関数
func extendLogLevelAttr(attr slog.Attr, key string) slog.Attr {
	// NOTE: LevelFatal slog.Level = 12 は ERROR+4 という文字列になる
	if attr.Key == slog.LevelKey && attr.Value.String() == "ERROR+4" {
		return slog.Attr{Key: key, Value: slog.StringValue("FATAL")}
	}
	return attr
}
