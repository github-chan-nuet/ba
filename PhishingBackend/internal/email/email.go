package email

import (
	"net/smtp"
)

type Email struct {
}

// https://zetcode.com/golang/email-smtp/
func SendEmail() error {
	from := "info@securaware.ch"

	user := "info" // os.Getenv("PHBA_SMTP_USER")
	password := "" // os.Getenv("PHBA_SMTP_PASSWORD")

	to := []string{
		"patrick.scheidegger@ost.ch",
	}

	addr := "mail.securaware.ch:25"
	host := "mail.securaware.ch"

	msg := []byte("From: info@securaware.ch\r\n" +
		"To: patrick.scheidegger@ost.ch\r\n" +
		"Subject: Test mail\r\n\r\n" +
		"Sehr geehrte Damen und Herren. Dies ist eine Test-E-Mail und kann vernachlässigt werden. Freundliche Grüsse, Patrick.\r\n")

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)
	return err
}
