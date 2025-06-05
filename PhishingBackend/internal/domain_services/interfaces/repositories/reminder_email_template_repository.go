package repositories

import "phishing_backend/internal/domain_model"

type ReminderEmailTemplateRepository interface {
	GetAll() (*[]domain_model.ReminderEmailTemplate, error)
}
