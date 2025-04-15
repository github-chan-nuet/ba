package db

import (
	"github.com/google/uuid"
	"log/slog"
	"phishing_backend/internal/domain/model"
)

var _ ExperienceRepository = (*ExperienceRepositoryImpl)(nil)

type ExperienceRepository interface {
	GetTotalExperience(userId uuid.UUID) (int, error)
}

type ExperienceRepositoryImpl struct {
}

func (e *ExperienceRepositoryImpl) GetTotalExperience(userId uuid.UUID) (int, error) {
	var count int64
	result := db.Model(&model.LessonCompletion{}).Where("UserFk = ?", userId).Count(&count)
	if result.Error != nil {
		slog.Error("Could not fetch experience from DB", "err", result.Error)
		return 0, result.Error
	}
	return 0, nil
}
