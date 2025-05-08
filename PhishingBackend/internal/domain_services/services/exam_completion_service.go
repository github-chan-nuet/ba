package services

import (
	"github.com/google/uuid"
	. "phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
)

var _ ExamCompletionService = (*ExamCompletionServiceImpl)(nil)

type ExamCompletionService interface {
	CompleteExam(userId, examId uuid.UUID, answers *[]QuestionCompletionDto) error
}

type ExamCompletionServiceImpl struct {
	Repo repositories.ExamRepository
}

func (e *ExamCompletionServiceImpl) CompleteExam(userId, examId uuid.UUID, answers *[]QuestionCompletionDto) error {
	exam, err := e.Repo.Get(examId)
	if err != nil {
		return err
	}

	return nil
}

// calculateScore calculates the score from [0%, 100%]
func (e *ExamCompletionServiceImpl) calculateScore(exam *Exam, answers *[]QuestionCompletionDto) int {
	scorePerQ := 100 / float32(len(exam.Questions))

	return 0
}

func (e *ExamCompletionServiceImpl) mapToUserQuestionResponse(exam *Exam, answers *[]QuestionCompletionDto) *[]userQuestionResponse {
	responses := make([]userQuestionResponse, 0, len(exam.Questions))
	for _, question := range exam.Questions {

	}
	return &responses
}

type userQuestionResponse struct {
	question    *ExamQuestion
	userAnswers []uuid.UUID
}
