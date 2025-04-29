package repositories

import (
	"errors"
	"github.com/google/uuid"
	"phishing_backend/internal/domain_model"
)

var LessonAlreadyCompleted = errors.New("lesson already completed")

type LessonCompletionRepository interface {
	Create(cc *domain_model.LessonCompletion) (int64, error)
	CountForUser(userId uuid.UUID) (int64, error)
	GetAllCompletedLessonsInAllCourses(userId uuid.UUID) ([]domain_model.LessonCompletion, error)
	GetLessonCompletionsOfCourseAndUser(userId, courseId uuid.UUID) ([]domain_model.LessonCompletion, error)
}
