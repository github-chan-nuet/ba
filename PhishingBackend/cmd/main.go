package main

import (
	"log/slog"
	"net/smtp"
	"os"
	"phishing_backend/internal/adapters"
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
		SmtpUser:   smtpUser,
		SmtpPw:     smtpPw,
		SmtpAddr:   smtpAddr,
		SmtpHost:   smtpHost,
		SendMailFn: smtp.SendMail,
	}
	emailRepository := persistence.EmailRepositoryImpl{}
	userRepository := persistence.UserRepositoryImpl{}
	phishingSimulationRepository := persistence.PhishingSimulationRepositoryImpl{}
	phishingEmailGenerationService := services.PhishingEmailGenerationServiceImpl{}
	phishingRunService := services.PhishingRunServiceImpl{
		EmailRepository:                &emailRepository,
		EmailSender:                    &emailSender,
		PhishingSimulationRepository:   &phishingSimulationRepository,
		PhishingEmailGenerationService: &phishingEmailGenerationService,
	}

	phishingOrchestrator := services.PhishingOrchestratorImpl{
		UserRepository:               &userRepository,
		PhishingSimulationRepository: &phishingSimulationRepository,
		PhishingRunService:           &phishingRunService,
	}
	phishingOrchestrator.StartPhishingRunGenerationJob()
	phishingOrchestrator.StartPhishingRunStregthDetectionJob()

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
