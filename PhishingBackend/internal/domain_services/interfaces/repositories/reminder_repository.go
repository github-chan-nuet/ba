package repositories

import domain "phishing_backend/internal/domain_model"

type ReminderRepository interface {
	GetAll() ([]domain.Reminder, error)
	SaveOrUpdate(reminder *domain.Reminder) error
}
