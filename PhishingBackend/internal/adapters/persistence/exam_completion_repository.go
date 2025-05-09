package persistence

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
)

var _ repositories.ExamCompletionRepository = (*ExamCompletionRepositoryImpl)(nil)

const uniqueExamCompletion = "unique_exam_completion_per_usr"

type ExamCompletionRepositoryImpl struct {
}

type ExamScore struct {
	Score int
}

func (r *ExamCompletionRepositoryImpl) GetScores(userId uuid.UUID) ([]int, error) {
	var scores []ExamScore
	result := db.Model(&domain_model.ExamCompletion{}).
		Select("score").
		Where("user_fk = ?", userId).
		Find(&scores)
	if result.Error != nil {
		slog.Error("Could not fetch scores in exam completion of users", "err", result.Error)
		return nil, result.Error
	}
	var intScores = make([]int, 0, len(scores))
	for _, score := range scores {
		intScores = append(intScores, score.Score)
	}
	return intScores, nil
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
