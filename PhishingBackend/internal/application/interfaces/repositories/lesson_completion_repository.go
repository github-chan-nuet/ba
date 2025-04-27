package repositories

import (
	"github.com/google/uuid"
	"phishing_backend/internal/domain"
)

type LessonCompletionRepository interface {
	Create(cc *domain.LessonCompletion) (int64, error)
	CountForUser(userId uuid.UUID) (int64, error)
}
