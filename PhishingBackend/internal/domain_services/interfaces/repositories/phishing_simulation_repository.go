package repositories

import (
	"phishing_backend/internal/domain_model"

	"github.com/google/uuid"
)

type PhishingSimulationRepository interface {
	Create(run *domain_model.PhishingSimulationRun) error
	Update(run *domain_model.PhishingSimulationRunPatch) error
	GetLatestRun(userId uuid.UUID) (*domain_model.PhishingSimulationRun, error)
	GetTemplates() ([]domain_model.PhishingSimulationContentTemplate, error)
	GetUserVulnerabilities(userId uuid.UUID) ([]domain_model.PhishingSimulationUserVulnerability, error)
	GetRecognitionFeatures() ([]domain_model.PhishingSimulationRecognitionFeature, error)
}
