package main

import (
	"log/slog"
	"os"
	"phishing_backend/internal/adapters"
	"phishing_backend/internal/adapters/presentation"
)

func main() {
	d := adapters.ResolveDependencies()
	go startReminderJob(d)
	presentation.SetupHttpServer(d)
}

func startReminderJob(d *adapters.Dependencies) {
	d.ReminderOrchestrator.ExecuteReminderJobAfterDurationEachDay()
}

func init() {
	setupDefaultLogger()
}

func setupDefaultLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
