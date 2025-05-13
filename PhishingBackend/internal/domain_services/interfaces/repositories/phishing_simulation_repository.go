package repositories

import (
	"github.com/google/uuid"
	"phishing_backend/internal/domain_model"
)

type PhishingSimulationRepository interface {
	Create(run *domain_model.PhishingSimulationRun) error
	Update(run *domain_model.PhishingSimulationRun) error
	GetLatestRun(userId uuid.UUID) (*domain_model.PhishingSimulationRun, error)
	GetTemplates() ([]domain_model.PhishingSimulationTemplate, error)
}
