package repositories

import "phishing_backend/internal/domain"

type LessonCompletionRepository interface {
	Create(cc *domain.LessonCompletion) (int64, error)
}
