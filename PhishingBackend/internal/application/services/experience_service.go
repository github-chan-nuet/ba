package services

import (
	"github.com/google/uuid"
	"phishing_backend/internal/application/interfaces/repositories"
)

var _ ExperienceService = (*ExperienceServiceImpl)(nil)

type UserExperience struct {
	TotalExperience int
	Level           int
}

type ExperienceService interface {
	Get(userId uuid.UUID) (*UserExperience, error)
}

type ExperienceServiceImpl struct {
	Repo repositories.LessonCompletionRepository
}

func (s *ExperienceServiceImpl) Get(userId uuid.UUID) (*UserExperience, error) {
	//TODO implement me
	panic("implement me")
}
