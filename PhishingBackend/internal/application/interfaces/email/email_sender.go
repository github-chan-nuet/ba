package email

type EmailSender interface {
	Send(content, recipient string) error
}
