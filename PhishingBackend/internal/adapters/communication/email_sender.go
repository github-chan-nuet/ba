package communication

import (
	"bytes"
	"log/slog"
	"mime"
	"mime/quotedprintable"
	"net/smtp"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/email"
)

var _ email.EmailSender = (*EmailSenderImpl)(nil)

const newLine = "\r\n"

type EmailSenderImpl struct {
	SmtpUser   string
	SmtpPw     string
	SmtpAddr   string
	SmtpHost   string
	SendMailFn func(addr string, a smtp.Auth, from string, to []string, msg []byte) error
}

func (e *EmailSenderImpl) Send(email *domain_model.Email) error {
	var msg bytes.Buffer

	to := []string{email.Recipient}

	encodedSubject := mime.QEncoding.Encode("utf-8", email.Subject)

	msg.WriteString("From: " + email.Sender + newLine)
	msg.WriteString("To: " + email.Recipient + newLine)
	msg.WriteString("Subject: " + encodedSubject + newLine)
	msg.WriteString("MIME-Version: 1.0" + newLine)
	msg.WriteString("Content-Type: text/html; charset=utf-8" + newLine)
	msg.WriteString("Content-Transfer-Encoding: quoted-printable" + newLine)
	msg.WriteString(newLine)

	qpWriter := quotedprintable.NewWriter(&msg)
	qpWriter.Write([]byte(email.Content))
	qpWriter.Close()

	auth := smtp.PlainAuth("", e.SmtpUser, e.SmtpPw, e.SmtpHost)

	err := e.SendMailFn(e.SmtpAddr, auth, email.Sender, to, msg.Bytes())
	if err != nil {
		slog.Error("Could not send out email", "recipient", email.Recipient, "subject", email.Subject, "error", err)
	}
	return err
}
