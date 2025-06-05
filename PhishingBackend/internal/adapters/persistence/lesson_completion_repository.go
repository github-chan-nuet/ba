package persistence

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"time"
)

var _ repositories.LessonCompletionRepository = (*LessonCompletionRepositoryImpl)(nil)

const uniqueLessonCompletion = "unique_lesson_completion_per_usr"

type LessonCompletionRepositoryImpl struct {
}

func (c *LessonCompletionRepositoryImpl) GetLatestLessonCompletions() (map[uuid.UUID]time.Time, error) {
	rows, err := db.Table("lesson_completion").
		Select("userfk as user, max(time) as last_completion_time").
		Group("userfk").
		Rows()
	if err != nil {
		slog.Error("Could not get all the latest lesson completions", "err", err)
		return nil, err
	}

	lasts := make(map[uuid.UUID]time.Time)
	for rows.Next() {
		var userId uuid.UUID
		var lastTime time.Time
		if err := rows.Scan(&userId, &lastTime); err != nil {
			slog.Error("Could not get userId and time of result", "err", err)
			return nil, err
		}
		lasts[userId] = lastTime
	}
	return lasts, nil
}

func (c *LessonCompletionRepositoryImpl) GetLessonCompletionsOfCourseAndUser(userId, courseId uuid.UUID) ([]domain_model.LessonCompletion, error) {
	var lessonCompletions []domain_model.LessonCompletion
	result := db.Where("user_fk = ? AND course_id = ?", userId, courseId).Find(&lessonCompletions)
	if result.Error != nil {
		slog.Error("Could not fetch lesson completions by user and courseId", "err", result.Error)
		return nil, result.Error
	}
	return lessonCompletions, nil
}

func (c *LessonCompletionRepositoryImpl) GetAllCompletedLessonsInAllCourses(userId uuid.UUID) ([]domain_model.LessonCompletion, error) {
	var lessonCompletions []domain_model.LessonCompletion
	result := db.Where("user_fk = ?", userId).Find(&lessonCompletions)
	if result.Error != nil {
		slog.Error("Could not fetch lesson completions by user", "err", result.Error)
		return nil, result.Error
	}
	return lessonCompletions, nil
}

func (c *LessonCompletionRepositoryImpl) CountForUser(userId uuid.UUID) (int, error) {
	var count int64
	result := db.Model(&domain_model.LessonCompletion{}).Where("user_fk = ?", userId).Count(&count)
	if result.Error != nil {
		slog.Error("Could not count lesson completions", "err", result.Error)
		return 0, result.Error
	}
	return int(count), nil // cast int64 down to int as int provides enough space
}

func (c *LessonCompletionRepositoryImpl) Create(cc *domain_model.LessonCompletion) (int, error) {
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
	return int(result.RowsAffected), nil
}
