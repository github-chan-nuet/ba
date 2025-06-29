package communication

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log/slog"
	"mime"
	"mime/quotedprintable"
	"net/smtp"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/email"
	"time"
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

	if email.Sender == "" {
		email.Sender = "Securaware <info@securaware.ch>"
	}

	// Generate Date header in RFC1123Z format (compatible with RFC5322)
	date := time.Now().UTC().Format(time.RFC1123Z)

	// Generate a Message-ID header
	// Format: <timestamp.randomString@securaware.ch>
	timestamp := time.Now().UnixNano()
	randomPart := make([]byte, 8)
	if _, err := rand.Read(randomPart); err != nil {
		// fallback if rand fails
		randomPart = []byte("fallback")
	}
	randomStr := base64.RawURLEncoding.EncodeToString(randomPart)
	messageID := fmt.Sprintf("<%d.%s@securaware.ch>", timestamp, randomStr)

	encodedSubject := mime.QEncoding.Encode("utf-8", email.Subject)

	msg.WriteString("From: " + email.Sender + newLine)
	msg.WriteString("To: " + email.Recipient + newLine)
	msg.WriteString("Subject: " + encodedSubject + newLine)
	msg.WriteString("Date: " + date + newLine)
	msg.WriteString("Message-ID: " + messageID + newLine)
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
