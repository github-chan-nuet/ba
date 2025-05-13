package repositories

import (
	"errors"
	"github.com/google/uuid"
	"phishing_backend/internal/domain_model"
)

var ErrExamAlreadyCompleted = errors.New("exam was already completed")

type ExamCompletionRepository interface {
	Save(exComp *domain_model.ExamCompletion) error
	GetScores(userId uuid.UUID) ([]int, error)
	GetCompletedExam(userId, examId uuid.UUID) (*domain_model.ExamCompletion, error)
}
