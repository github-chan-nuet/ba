package services

import (
	"errors"
	"github.com/google/uuid"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"time"
)

type IsNewEntry bool

var _ LessonCompletionService = (*LessonCompletionServiceImpl)(nil)

type LessonCompletionService interface {
	Create(courseId, lessonId, userId uuid.UUID) (IsNewEntry, error)
}

type LessonCompletionServiceImpl struct {
	Repo repositories.LessonCompletionRepository
}

func (c *LessonCompletionServiceImpl) Create(courseId, lessonId, userId uuid.UUID) (IsNewEntry, error) {
	lc := domain_model.LessonCompletion{
		ID:       uuid.New(),
		LessonId: lessonId,
		CourseId: courseId,
		UserFk:   userId,
		Time:     time.Now().UTC(),
	}
	_, err := c.Repo.Create(&lc)
	if errors.Is(err, repositories.LessonAlreadyCompleted) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
