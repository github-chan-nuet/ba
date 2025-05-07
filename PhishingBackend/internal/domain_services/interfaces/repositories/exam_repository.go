package repositories

import (
	"github.com/google/uuid"
	"phishing_backend/internal/domain_model"
)

type ExamRepository interface {
	Get(examId uuid.UUID) (*domain_model.Exam, error)
}
