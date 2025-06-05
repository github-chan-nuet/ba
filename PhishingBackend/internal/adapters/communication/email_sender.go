package communication

import (
	"log/slog"
	"net/smtp"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/email"
)

var _ email.EmailSender = (*EmailSenderImpl)(nil)

const newLine = "\r\n"

type EmailSenderImpl struct {
	SmtpUser string
	SmtpPw   string
	SmtpAddr string
	SmtpHost string
}

func (e *EmailSenderImpl) Send(email *domain_model.Email) error {
	from := "info@securaware.ch"
	to := []string{email.Subject}

	msg := []byte(
		"From: " + from + newLine +
			"To: " + email.Recipient + newLine +
			"Subject: " + email.Subject + newLine + newLine +
			email.Content + newLine)

	auth := smtp.PlainAuth("", e.SmtpUser, e.SmtpPw, e.SmtpHost)

	err := smtp.SendMail(e.SmtpAddr, auth, from, to, msg)
	if err != nil {
		slog.Error("Could not send out email", "recipient", email.Recipient, "subject", email.Subject, "error", err)
	}
	return err
}

//func (e *EmailSenderImpl) Send(email *domain_model.Email) error {
//	var msg bytes.Buffer
//
//	from := "info@securaware.ch"
//	to := []string{email.Recipient}
//
//	encodedSubject := mime.QEncoding.Encode("utf-8", email.Subject)
//
//	msg.WriteString("From: " + from + newLine)
//	msg.WriteString("To: " + email.Recipient + newLine)
//	msg.WriteString("Subject: " + encodedSubject + newLine)
//	msg.WriteString("MIME-Version: 1.0" + newLine)
//	msg.WriteString("Content-Type: text/html; charset=utf-8" + newLine)
//	msg.WriteString("Content-Transfer-Encoding: quoted-printable" + newLine)
//	msg.WriteString(newLine)
//
//	qpWriter := quotedprintable.NewWriter(&msg)
//	qpWriter.Write([]byte(email.Content))
//	qpWriter.Close()
//
//	auth := smtp.PlainAuth("", e.SmtpUser, e.SmtpPw, e.SmtpHost)
//
//	err := smtp.SendMail(e.SmtpAddr, auth, from, to, msg.Bytes())
//	return err
//}
