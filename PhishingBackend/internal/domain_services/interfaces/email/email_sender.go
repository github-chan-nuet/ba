package email

import "phishing_backend/internal/domain_model"

type EmailSender interface {
	Send(email *domain_model.Email) error
}
