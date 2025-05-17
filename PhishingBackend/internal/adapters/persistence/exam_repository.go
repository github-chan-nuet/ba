package persistence

import (
	"github.com/google/uuid"
	"log/slog"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
)

var _ repositories.ExamRepository = (*ExamRepositoryImpl)(nil)

type ExamRepositoryImpl struct {
}

func (e *ExamRepositoryImpl) GetExamIds() ([]uuid.UUID, error) {
	var exams []domain_model.Exam
	result := db.Model(&domain_model.Exam{}).Find(&exams)
	if result.Error != nil {
		slog.Error("Could not get exam ids", "err", result.Error)
		return nil, result.Error
	}
	examIds := make([]uuid.UUID, 0, len(exams))
	for _, exam := range exams {
		examIds = append(examIds, exam.ID)
	}
	return examIds, nil
}

func (e *ExamRepositoryImpl) Get(examId uuid.UUID) (*domain_model.Exam, error) {
	var exam domain_model.Exam
	result := db.Model(&domain_model.Exam{}).Preload("Questions.Answers").Where("ID = ?", examId).Find(&exam)
	if result.Error != nil {
		slog.Error("Could not get exam by id", "err", result.Error)
		return nil, result.Error
	}
	return &exam, nil
}
