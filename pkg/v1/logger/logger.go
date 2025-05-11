package logger

import (
	"log/slog"
	"os"
	"sync"

	slogformatter "github.com/samber/slog-formatter"
)

var defaultLogger *slog.Logger
var initLoggerOnce sync.Once

// InitializeLogger initializes the logger with default settings
// and sets the default logger to be used throughout the application.
func InitializeLogger() *slog.Logger {
	initLoggerOnce.Do(func() {
		defaultLogger = slog.New(
			slogformatter.NewFormatterHandler(
				slogformatter.ErrorFormatter("error"),
			)(
				slog.NewJSONHandler(os.Stderr, nil),
			),
		)
	})

	return defaultLogger
}
