package repositories

import (
	"errors"
	"github.com/google/uuid"
	"phishing_backend/internal/domain_model"
	"time"
)

var LessonAlreadyCompleted = errors.New("lesson already completed")

type LessonCompletionRepository interface {
	Create(cc *domain_model.LessonCompletion) (int, error)
	CountForUser(userId uuid.UUID) (int, error)
	GetAllCompletedLessonsInAllCourses(userId uuid.UUID) ([]domain_model.LessonCompletion, error)
	GetLessonCompletionsOfCourseAndUser(userId, courseId uuid.UUID) ([]domain_model.LessonCompletion, error)
	GetLatestLessonCompletions() (map[uuid.UUID]time.Time, error)
}
