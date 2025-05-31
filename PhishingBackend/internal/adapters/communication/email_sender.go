package communication

import (
	"net/smtp"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/email"
)

var _ email.EmailSender = (*EmailSenderImpl)(nil)

const newLine = "\r\n"

//smtpUser                   = os.Getenv("PHBA_SMTP_USER")
//smtpPw                     = os.Getenv("PHBA_SMTP_PASSWORD")
//smtpAddr                   = os.Getenv("PHBA_SMTP_ADDR")
//smtpHost                   = os.Getenv("PHBA_SMTP_HOST")

type EmailSenderImpl struct {
	SmtpUser string
	SmtpPw   string
	SmtpAddr string
	SmtpHost string
}

// https://zetcode.com/golang/email-smtp/
func (e *EmailSenderImpl) Send(email *domain_model.Email) error {
	from := "info@securaware.ch"
	to := []string{email.Recipient}

	msg := []byte(
		"From: " + from + newLine +
			"To: " + email.Recipient + newLine +
			"Subject: " + email.Subject + newLine + newLine +
			email.Content + newLine)

	auth := smtp.PlainAuth("", e.SmtpUser, e.SmtpPw, e.SmtpHost)

	err := smtp.SendMail(e.SmtpAddr, auth, from, to, msg)
	return err
}
