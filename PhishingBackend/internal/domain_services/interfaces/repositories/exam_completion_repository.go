package repositories

import (
	"errors"
	"phishing_backend/internal/domain_model"
)

var ErrExamAlreadyCompleted = errors.New("exam was already completed")

type ExamCompletionRepository interface {
	Save(exComp *domain_model.ExamCompletion) error
}
