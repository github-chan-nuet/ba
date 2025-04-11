package db

import (
	"log/slog"
	"phishing_backend/internal/domain/model"
)

var _ CourseCompletionRepository = (*CourseCompletionRepositoryImpl)(nil)

type CourseCompletionRepository interface {
	Create(cc *model.CourseCompletion) (int64, error)
}

type CourseCompletionRepositoryImpl struct {
}

func (c *CourseCompletionRepositoryImpl) Create(cc *model.CourseCompletion) (int64, error) {
	result := db.Create(cc)
	if result.Error != nil {
		slog.Error("Could not create course completion", "err", result.Error)
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
