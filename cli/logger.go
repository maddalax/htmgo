package main

import (
	"log/slog"
	"os"
	"strings"
)

func getLogLevel() slog.Level {
	// Get the log level from the environment variable
	logLevel := os.Getenv("LOG_LEVEL")
	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		// Default to INFO if no valid log level is set
		return slog.LevelInfo
	}
}
