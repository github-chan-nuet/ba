package persistence

import (
	"log/slog"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
)

var _ repositories.ReminderEmailTemplateRepository = (*ReminderEmailTemplateRepositoryImpl)(nil)

type ReminderEmailTemplateRepositoryImpl struct{}

func (r *ReminderEmailTemplateRepositoryImpl) GetAll() (*[]domain_model.ReminderEmailTemplate, error) {
	var templates []domain_model.ReminderEmailTemplate
	result := db.Find(&templates)
	if result.Error != nil {
		slog.Error("Could not get all templates", "err", result.Error)
		return nil, result.Error
	}
	return &templates, nil
}
