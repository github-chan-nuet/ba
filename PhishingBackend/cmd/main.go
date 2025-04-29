package main

import (
	"log/slog"
	"os"
	"phishing_backend/internal/infrastructure/presentation"
)

func main() {
	presentation.SetupHttpServer()
}

func init() {
	setupDefaultLogger()
}

func setupDefaultLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
