package experience

import (
	"github.com/google/uuid"
	"phishing_backend/internal/domain/db"
)

var _ Service = (*ServiceImpl)(nil)

type UserExperience struct {
	TotalExperience int
	Level           int
}

type Service interface {
	Get(userId uuid.UUID) (*UserExperience, error)
}

type ServiceImpl struct {
	Repo db.LessonCompletionRepository
}

func (s *ServiceImpl) Get(userId uuid.UUID) (*UserExperience, error) {
	//TODO implement me
	panic("implement me")
}
