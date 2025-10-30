package log

import (
	"log/slog"
)

var (
	Log Logger = NewSLogger(slog.Default())
)

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
	Debug(msg string, args ...any)
}

type SLogger struct {
	logger *slog.Logger
}

func (S *SLogger) Info(msg string, args ...any) {
	S.logger.Info(msg, toAttrs(args))
}

func (S *SLogger) Error(msg string, args ...any) {
	S.logger.Error(msg, toAttrs(args))
}

func (S *SLogger) Debug(msg string, args ...any) {
	S.logger.Debug(msg, toAttrs(args))
}

func NewSLogger(logger *slog.Logger) *SLogger {
	return &SLogger{logger: logger}
}

func toAttrs(pairs []any) []slog.Attr {
	var attrs []slog.Attr
	for i := 0; i < len(pairs); i += 2 {
		if i+1 < len(pairs) {
			key, ok := pairs[i].(string)
			if !ok {
				continue // 跳过无效的 key
			}
			attrs = append(attrs, slog.Any(key, pairs[i+1]))
		} else {
			// 单数个参数，最后一个没有 value
			attrs = append(attrs, slog.Any("unknown", pairs[i]))
		}
	}
	return attrs
}
