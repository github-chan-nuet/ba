package communication

import (
	"bytes"
	"mime"
	"net/smtp"
	"phishing_backend/internal/domain_model"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type sendMailParams struct {
	addr string
	a    smtp.Auth
	from string
	to   []string
	msg  []byte
}

var sendMailArgs sendMailParams

func createEmailSenderImpl() *EmailSenderImpl {
	return &EmailSenderImpl{
		SmtpUser: "user",
		SmtpPw:   "pw",
		SmtpAddr: "addr",
		SmtpHost: "host",
		NowFn:    getFixedDate,
		MsgIDFn:  getFixedMessageId,
		SendMailFn: func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
			sendMailArgs = sendMailParams{
				addr: addr,
				a:    a,
				from: from,
				to:   to,
				msg:  msg,
			}
			return nil
		},
	}
}

func getFixedDate() time.Time {
	return time.Date(2025, 5, 12, 13, 45, 23, 0, time.UTC)
}

func getFixedMessageId() string {
	return "<fixed-message-id@domain>"
}

func getWantMessage(email *domain_model.Email, content string) []byte {
	var msg bytes.Buffer
	msg.WriteString("From: " + email.Sender + newLine)
	msg.WriteString("To: " + email.Recipient + newLine)
	encodedSubject := mime.QEncoding.Encode("utf-8", email.Subject)
	msg.WriteString("Subject: " + encodedSubject + newLine)
	msg.WriteString("Date: " + getFixedDate().Format(time.RFC1123Z) + newLine)
	msg.WriteString("Message-ID: " + getFixedMessageId() + newLine)
	msg.WriteString("MIME-Version: 1.0" + newLine)
	msg.WriteString("Content-Type: text/html; charset=utf-8" + newLine)
	msg.WriteString("Content-Transfer-Encoding: quoted-printable" + newLine)
	msg.WriteString(newLine)
	msg.WriteString(content)
	return msg.Bytes()
}

func TestSend(t *testing.T) {
	// given
	sut := createEmailSenderImpl()
	email := domain_model.Email{
		ID:        uuid.UUID{},
		Sender:    "a@a",
		Recipient: "b@b",
		Subject:   "subject",
		Content:   "äöüÄÖÜ",
	}

	// when
	sut.Send(&email)

	// then
	assert.Equal(t, sut.SmtpAddr, sendMailArgs.addr)
	wantAuth := smtp.PlainAuth("", sut.SmtpUser, sut.SmtpPw, sut.SmtpHost)
	assert.Equal(t, wantAuth, sendMailArgs.a)
	assert.Equal(t, email.Sender, sendMailArgs.from)
	assert.Equal(t, []string{email.Recipient}, sendMailArgs.to)
}

func TestSendContent(t *testing.T) {
	// given
	sut := createEmailSenderImpl()
	email := domain_model.Email{
		ID:        uuid.UUID{},
		Sender:    "a@a",
		Recipient: "b@b",
		Subject:   "subject",
	}

	tests := []struct {
		name    string
		content string
		want    string
	}{{
		name:    "Special characters",
		content: "äöüÄÖÜ",
		want:    "=C3=A4=C3=B6=C3=BC=C3=84=C3=96=C3=9C",
	}, {
		name:    "75 chars",
		content: strings.Repeat("a", 75),
		want:    strings.Repeat("a", 75),
	}, {
		name:    "76 chars new line",
		content: strings.Repeat("a", 76),
		want:    strings.Repeat("a", 75) + "=" + newLine + "a",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email.Content = tt.content

			// when
			sut.Send(&email)

			// then
			wantMessage := getWantMessage(&email, tt.want)
			assert.Equal(t, string(wantMessage), string(sendMailArgs.msg))
		})
	}

}
