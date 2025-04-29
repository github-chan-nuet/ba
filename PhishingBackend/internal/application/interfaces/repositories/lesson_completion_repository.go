package repositories

import (
	"errors"
	"github.com/google/uuid"
	"phishing_backend/internal/domain"
)

var LessonAlreadyCompleted = errors.New("lesson already completed")

type LessonCompletionRepository interface {
	Create(cc *domain.LessonCompletion) (int64, error)
	CountForUser(userId uuid.UUID) (int64, error)
}
