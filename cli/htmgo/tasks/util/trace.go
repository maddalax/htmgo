package util

import (
	"log/slog"
	"time"
)

func Trace(name string, cb func() any) any {
	now := time.Now()
	result := cb()
	slog.Debug("trace", slog.String("name", name), slog.Duration("duration", time.Since(now)))
	return result
}
