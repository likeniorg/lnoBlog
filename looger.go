package main

import (
	"log/slog"
	"os"
)

var looger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func logInfo(msg string, args ...any) {
	looger.Info(msg, args...)

}

func logWarn(msg string, args ...any) {
	looger.Warn(msg, args...)

}

func logError(msg string, args ...any) {
	looger.Error(msg, args...)

}
