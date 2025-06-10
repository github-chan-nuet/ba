package persistence

import (
	"errors"
	"log/slog"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
)

var _ repositories.EmailRepository = (*EmailRepositoryImpl)(nil)

type EmailRepositoryImpl struct {
}

func (r *EmailRepositoryImpl) Create(email *domain_model.Email) error {
	result := db.Create(email)
	if result.Error != nil {
		slog.Error("Could not save email record")
	}
	return result.Error
}

func (r *EmailRepositoryImpl) Update(emailPatch *domain_model.EmailPatch) error {
	var existing domain_model.Email
	if err := db.First(&existing, emailPatch.ID).Error; err != nil {
		return errors.New("Email record not found")
	}

	updates := map[string]interface{}{}

	if emailPatch.SentAt != nil {
		if existing.SentAt != nil {
			return errors.New("SentAt is already set")
		}
		updates["sent_at"] = *emailPatch.SentAt
	}

	if emailPatch.ClickedAt != nil {
		if existing.ClickedAt != nil {
			return errors.New("ClickedAt is already set")
		}
		updates["clicked_at"] = *emailPatch.ClickedAt
	}

	if len(updates) > 0 {
		if err := db.Model(&domain_model.Email{}).
			Where("id = ?", emailPatch.ID).
			Updates(updates).Error; err != nil {
			slog.Error("Could not update Email record")
		}
	}
	return nil
}
