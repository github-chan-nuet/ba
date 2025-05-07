package services

import (
	"github.com/google/uuid"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
)

var _ ExamCompletionService = (*ExamCompletionServiceImpl)(nil)

type ExamCompletionService interface {
	CompleteExam(userId, examId uuid.UUID, answers *[]domain_model.QuestionCompletionDto) error
}

type ExamCompletionServiceImpl struct {
	Repo repositories.ExamRepository
}

func (e *ExamCompletionServiceImpl) CompleteExam(userId, examId uuid.UUID, answers *[]domain_model.QuestionCompletionDto) error {
	exam, err := e.Repo.Get(examId)
	if err != nil {
		return err
	}

	return nil
}
