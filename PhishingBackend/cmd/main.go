package main

import (
	"log/slog"
	"os"
	"phishing_backend/internal/adapters"
	"phishing_backend/internal/adapters/presentation"
)

func main() {
	d := adapters.ResolveDependencies()
	startReminderJob(d)
	startPhishingJobs(d)
	presentation.SetupHttpServer(d)
}

func startReminderJob(d *adapters.Dependencies) {
	go d.ReminderOrchestrator.ExecuteReminderJobAfterDurationEachDay()
}

func startPhishingJobs(d *adapters.Dependencies) {
	go d.PhishingOrchestrator.StartPhishingRunGenerationJob()
	go d.PhishingOrchestrator.StartPhishingRunStregthDetectionJob()
}

func init() {
	setupDefaultLogger()
}

func setupDefaultLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
