package email

import (
	"net/smtp"
	"os"
)

type Email struct {
}

func SendEmail() error {
	from := "john.doe@example.com"

	user := os.Getenv("PHBA_SMTP_USER")
	password := os.Getenv("PHBA_SMTP_PASSWORD")

	to := []string{
		"roger.roe@example.com",
	}

	addr := "smtp.mailtrap.io:2525"
	host := "smtp.mailtrap.io"

	msg := []byte("From: john.doe@example.com\r\n" +
		"To: roger.roe@example.com\r\n" +
		"Subject: Test mail\r\n\r\n" +
		"Email body\r\n")

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)
	return err
}
