package services

import (
	"os"
	"phishing_backend/internal/adapters/communication"
	"phishing_backend/internal/domain_model"
	"time"
)

var _ ReminderOrchestrator = (*ReminderOrchestratorImpl)(nil)

type ReminderOrchestrator interface {
	StartReminderJob(d time.Duration)
}

type ReminderOrchestratorImpl struct {
}

func (r *ReminderOrchestratorImpl) StartReminderJob(d time.Duration) {
	go StartCronJob(d, r.sendReminder)
}

func (r *ReminderOrchestratorImpl) sendReminder(utc time.Time) {
	smtpUser := os.Getenv("PHBA_SMTP_USER")
	smtpPw := os.Getenv("PHBA_SMTP_PASSWORD")
	smtpAddr := os.Getenv("PHBA_SMTP_ADDR")
	smtpHost := os.Getenv("PHBA_SMTP_HOST")

	mailer := communication.EmailSenderImpl{
		SmtpUser: smtpUser,
		SmtpPw:   smtpPw,
		SmtpAddr: smtpAddr,
		SmtpHost: smtpHost,
	}

	reminder := domain_model.Email{
		Content:   "Dies ist ein Test-Reminder",
		Recipient: "mischa.binder@stafag.ch",
		Subject:   "Test-Reminder | Securaware",
	}

	mailer.Send(&reminder)
}
