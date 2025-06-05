package main

import (
	"log/slog"
	"os"
	"phishing_backend/internal/adapters/communication"
	"phishing_backend/internal/adapters/persistence"
	"phishing_backend/internal/adapters/presentation"
	"phishing_backend/internal/domain_services/services"
)

func main() {
	smtpUser := os.Getenv("PHBA_SMTP_USER")
	smtpPw := os.Getenv("PHBA_SMTP_PASSWORD")
	smtpAddr := os.Getenv("PHBA_SMTP_ADDR")
	smtpHost := os.Getenv("PHBA_SMTP_HOST")

	emailSender := communication.EmailSenderImpl{
		SmtpUser: smtpUser,
		SmtpPw:   smtpPw,
		SmtpAddr: smtpAddr,
		SmtpHost: smtpHost,
	}
	userRepository := persistence.UserRepositoryImpl{}
	phishingSimulationRepository := persistence.PhishingSimulationRepositoryImpl{}
	phishingEmailGenerationService := services.PhishingEmailGenerationServiceImpl{}
	phishingRunService := services.PhishingRunServiceImpl{
		EmailSender:                    &emailSender,
		PhishingSimulationRepository:   &phishingSimulationRepository,
		PhishingEmailGenerationService: &phishingEmailGenerationService,
	}

	phishingOrchestrator := services.PhishingOrchestratorImpl{
		UserRepository:               &userRepository,
		PhishingSimulationRepository: &phishingSimulationRepository,
		PhishingRunService:           &phishingRunService,
	}
	phishingOrchestrator.StartPhishingJob()

	presentation.SetupHttpServer()
}

func init() {
	setupDefaultLogger()
}

func setupDefaultLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
