package services

import (
	"github.com/google/uuid"
	"math"
	"phishing_backend/internal/application/interfaces/repositories"
	"phishing_backend/internal/domain"
)

var _ ExperienceService = (*ExperienceServiceImpl)(nil)

type ExperienceService interface {
	GetEntireExperience(userId uuid.UUID) (*domain.UserExperience, error)
	GetExperienceGain(userId uuid.UUID, gain int) (*domain.ExperienceGain, error)
	calcLevel(totalExperience int) int
}

type ExperienceServiceImpl struct {
	Repo repositories.LessonCompletionRepository
}

func (s *ExperienceServiceImpl) GetExperienceGain(userId uuid.UUID, gain int) (*domain.ExperienceGain, error) {
	exp, err := s.GetEntireExperience(userId)
	if err != nil {
		return nil, err
	}
	previousLvl := s.calcLevel(exp.TotalExperience - gain)
	expGain := domain.ExperienceGain{
		NewExperienceGained: gain,
		TotalExperience:     exp.TotalExperience,
	}
	if previousLvl < exp.Level {
		expGain.NewLevel = &exp.Level
	}
	return &expGain, nil
}

func (s *ExperienceServiceImpl) GetEntireExperience(userId uuid.UUID) (*domain.UserExperience, error) {
	numLessons, err := s.Repo.CountForUser(userId)
	if err != nil {
		return nil, err
	}
	totExp := int(numLessons * domain.LessonCompletionGain)
	level := s.calcLevel(totExp)
	userExp := domain.UserExperience{TotalExperience: totExp, Level: level}
	return &userExp, nil
}

func (s *ExperienceServiceImpl) calcLevel(totalExperience int) int {
	// level is calculated by the formula: 1 + ln(x/200 + 1) / ln(1.5)
	// Desmos: 1\ +\ \ln\left(\frac{x}{200}+1\right)\ \frac{1}{\ln\left(1.5\right)}
	levelAsFloat := float64(1) + (math.Log((float64(totalExperience)/200)+1) / math.Log(1.5))
	lvl := int(math.Floor(levelAsFloat))
	return lvl
}
