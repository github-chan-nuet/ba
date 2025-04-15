package db

import (
	"log/slog"
	"phishing_backend/internal/domain/model"
)

var _ LessonCompletionRepository = (*LessonCompletionRepositoryImpl)(nil)

type LessonCompletionRepository interface {
	Create(cc *model.LessonCompletion) (int64, error)
}

type LessonCompletionRepositoryImpl struct {
}

func (c *LessonCompletionRepositoryImpl) Create(cc *model.LessonCompletion) (int64, error) {
	result := db.Create(cc)
	if result.Error != nil {
		slog.Error("Could not create course completion", "err", result.Error)
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
