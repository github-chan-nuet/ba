package persistence

import (
	"log/slog"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
)

var _ repositories.ReminderRepository = (*ReminderRepositoryImpl)(nil)

type ReminderRepositoryImpl struct {
}

func (r *ReminderRepositoryImpl) GetAll() ([]domain_model.Reminder, error) {
	var reminder []domain_model.Reminder
	result := db.Find(&reminder)
	if result.Error != nil {
		slog.Error("Could not fetch reminders", "err", result.Error)
		return nil, result.Error
	}
	return reminder, nil
}
