package courseCompletion

import (
	"phishing_backend/internal/domain/db"
	"phishing_backend/internal/domain/model"
)

type IsNewEntry bool

var _ Service = (*ServiceImpl)(nil)

type Service interface {
	Create(cc *model.CourseCompletion) (IsNewEntry, error)
}

type ServiceImpl struct {
	Repo db.CourseCompletionRepository
}

func (c ServiceImpl) Create(cc *model.CourseCompletion) (IsNewEntry, error) {
	count, err := c.Repo.Create(cc)
	if err != nil {
		return false, err
	}
	return count != 0, nil
}
