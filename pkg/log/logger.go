package log

import (
	"log/slog"
	"os"
)

var (
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	Info   = logger.Info
	Error  = logger.Error
	Fatal  = func(msg string) {
		logger.Error(msg)
		os.Exit(1)
	}
)
