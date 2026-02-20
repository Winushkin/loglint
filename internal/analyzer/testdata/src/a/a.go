package a

import (
	"log/slog"
)

func test() {
	slog.Debug("Starting server") // want "log message must start with lowercase letter"
	slog.Info("запуск")           // want "log message must be in English only"
	slog.Warn("!dadasdad!")       // want "log message contains special characters"
	slog.Error("password")        // want "log message may contain sensitive data"
}
