package persistance

import (
	"log/slog"
	"phishing_backend/internal/application/interfaces/repositories"
	"phishing_backend/internal/domain"
)

var _ repositories.LessonCompletionRepository = (*LessonCompletionRepositoryImpl)(nil)

type LessonCompletionRepositoryImpl struct {
}

func (c *LessonCompletionRepositoryImpl) Create(cc *domain.LessonCompletion) (int64, error) {
	result := db.Create(cc)
	if result.Error != nil {
		slog.Error("Could not create course completion", "err", result.Error)
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
