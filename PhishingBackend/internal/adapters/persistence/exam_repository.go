package persistence

import (
	"log/slog"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"

	"github.com/google/uuid"
)

var _ repositories.ExamRepository = (*ExamRepositoryImpl)(nil)

type ExamRepositoryImpl struct {
}

func (e *ExamRepositoryImpl) GetAll() (*[]domain_model.Exam, error) {
	var exams []domain_model.Exam
	result := db.Model(&domain_model.Exam{}).Preload("Questions.Answers").Find(&exams)
	if result.Error != nil {
		slog.Error("Coud not get exams", "err", result.Error)
		return nil, result.Error
	}
	return &exams, nil
}

func (e *ExamRepositoryImpl) Get(examId uuid.UUID) (*domain_model.Exam, error) {
	var exam domain_model.Exam
	result := db.Model(&domain_model.Exam{}).
		Preload("Questions.Answers").
		Where("ID = ?", examId).
		First(&exam)
	if result.Error != nil {
		slog.Error("Could not get exam by id", "err", result.Error)
		return nil, result.Error
	}
	return &exam, nil
}
