package persistence

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
)

var _ repositories.ExamCompletionRepository = (*ExamCompletionRepositoryImpl)(nil)

const uniqueExamCompletion = "unique_exam_completion_per_usr"

type ExamCompletionRepositoryImpl struct {
}

func (r *ExamCompletionRepositoryImpl) Save(exComp *domain_model.ExamCompletion) error {
	result := db.Save(exComp)
	if result.Error != nil {
		var e *pgconn.PgError
		if errors.As(result.Error, &e) {
			if e.Code == "23505" && e.ConstraintName == uniqueExamCompletion {
				return repositories.ErrExamAlreadyCompleted
			}
		}
		slog.Error("Could not save exam completion", "err", result.Error)
	}
	return result.Error
}
