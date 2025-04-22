package communication

import (
	"net/smtp"
	"phishing_backend/internal/application/interfaces/email"
)

var _ email.EmailSender = (*EmailSenderImpl)(nil)

type EmailSenderImpl struct{}

// https://zetcode.com/golang/email-smtp/
func (e *EmailSenderImpl) Send(content, recipient string) error {
	from := "info@securaware.ch"

	user := "info" // os.Getenv("PHBA_SMTP_USER")
	password := "" // os.Getenv("PHBA_SMTP_PASSWORD")

	to := []string{
		"patrick.scheidegger@ost.ch",
	}

	addr := "mail.securaware.ch:25"
	host := "mail.securaware.ch"

	msg := []byte("From: info@securaware.ch\r\n" +
		"To: " + recipient + "\r\n" +
		"Subject: Test mail\r\n\r\n" +
		content + "\r\n")

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)
	return err
}
