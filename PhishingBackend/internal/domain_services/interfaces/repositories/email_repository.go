package repositories

import "phishing_backend/internal/domain_model"

type EmailRepository interface {
	Create(email *domain_model.Email) error
	Update(emailPatch *domain_model.EmailPatch) error
}
