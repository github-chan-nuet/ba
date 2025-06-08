package repositories

import "phishing_backend/internal/domain_model"

type ReminderRepository interface {
	GetAll() ([]domain_model.Reminder, error)
	SaveOrUpdate(reminder *domain_model.Reminder) error
}
