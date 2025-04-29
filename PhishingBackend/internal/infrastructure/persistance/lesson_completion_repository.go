package persistance

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
	"phishing_backend/internal/application/interfaces/repositories"
	"phishing_backend/internal/domain"
)

var _ repositories.LessonCompletionRepository = (*LessonCompletionRepositoryImpl)(nil)

const uniqueLessonCompletion = "unique_lesson_completion_per_usr"

type LessonCompletionRepositoryImpl struct {
}

func (c *LessonCompletionRepositoryImpl) CountForUser(userId uuid.UUID) (int64, error) {
	var count int64
	result := db.Model(&domain.LessonCompletion{}).Where("user_fk = ?", userId).Count(&count)
	if result.Error != nil {
		slog.Error("Could not count lesson completions", "err", result.Error)
		return 0, result.Error
	}
	return count, nil
}

func (c *LessonCompletionRepositoryImpl) Create(cc *domain.LessonCompletion) (int64, error) {
	result := db.Create(cc)
	if result.Error != nil {
		var e *pgconn.PgError
		if errors.As(result.Error, &e) {
			if e.Code == "23505" && e.ConstraintName == uniqueLessonCompletion {
				return 0, repositories.LessonAlreadyCompleted
			}
		}
		slog.Error("Could not create lesson completion", "err", result.Error)
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
