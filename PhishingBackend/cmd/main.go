package main

import (
	"log/slog"
	"os"
	"phishing_backend/internal/adapters/presentation"
	"phishing_backend/internal/domain_services/services"
	"time"
)

func main() {
	phishingOrchestrator := services.PhishingOrchestratorImpl{}
	phishingOrchestrator.StartPhishingJob(2 * time.Minute)

	presentation.SetupHttpServer()
}

func init() {
	setupDefaultLogger()
}

func setupDefaultLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
