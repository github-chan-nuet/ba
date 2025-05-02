package communication

import (
	"net/smtp"
	"os"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/email"
)

var (
	_        email.EmailSender = (*EmailSenderImpl)(nil)
	smtpUser                   = os.Getenv("PHBA_SMTP_USER")
	smtpPw                     = os.Getenv("PHBA_SMTP_PASSWORD")
	smtpAddr                   = os.Getenv("PHBA_SMTP_ADDR")
	smtpHost                   = os.Getenv("PHBA_SMTP_HOST")
)

const newLine = "\r\n"

type EmailSenderImpl struct{}

// https://zetcode.com/golang/email-smtp/
func (e *EmailSenderImpl) Send(email *domain_model.Email) error {
	from := "info@securaware.ch"
	to := []string{email.Subject}

	msg := []byte(
		"From: " + from + newLine +
			"To: " + email.Recipient + newLine +
			"Subject: " + email.Subject + newLine + newLine +
			email.Content + newLine)

	auth := smtp.PlainAuth("", smtpUser, smtpPw, smtpHost)

	err := smtp.SendMail(smtpAddr, auth, from, to, msg)
	return err
}
