package persistance

import (
	"github.com/google/uuid"
	"log/slog"
	"phishing_backend/internal/application/interfaces/repositories"
	"phishing_backend/internal/domain"
)

var _ repositories.LessonCompletionRepository = (*LessonCompletionRepositoryImpl)(nil)

type LessonCompletionRepositoryImpl struct {
}

func (c *LessonCompletionRepositoryImpl) CountForUser(userId uuid.UUID) (int64, error) {
	var count int64
	result := db.Where("User = ?", userId).Count(&count)
	if result.Error != nil {
		slog.Error("Could not count lesson completions", "err", result.Error)
		return 0, result.Error
	}
	return count, nil
}

func (c *LessonCompletionRepositoryImpl) Create(cc *domain.LessonCompletion) (int64, error) {
	result := db.Create(cc)
	if result.Error != nil {
		slog.Error("Could not create lesson completion", "err", result.Error)
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
