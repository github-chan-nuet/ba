package repositories

import (
	"phishing_backend/internal/domain_model"

	"github.com/google/uuid"
)

type ExamRepository interface {
	Get(examId uuid.UUID) (*domain_model.Exam, error)
	GetAll() (*[]domain_model.Exam, error)
}
