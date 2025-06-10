package repositories

import (
	"errors"
	"phishing_backend/internal/domain_model"

	"github.com/google/uuid"
)

var ErrUserVulnAlreadyExists = errors.New("User Vulnerability already exists")

type PhishingSimulationRepository interface {
	Create(run *domain_model.PhishingSimulationRun) error
	Update(run *domain_model.PhishingSimulationRunPatch) error
	GetRun(runId uuid.UUID) (*domain_model.PhishingSimulationRun, error)
	GetLatestRun(userId uuid.UUID) (*domain_model.PhishingSimulationRun, error)
	GetUnprocessedRuns() ([]domain_model.PhishingSimulationRun, error)
	GetTemplates() ([]domain_model.PhishingSimulationContentTemplate, error)
	CreateUserVulnerability(vulnerability *domain_model.PhishingSimulationUserVulnerability) error
	GetUserVulnerabilities(userId uuid.UUID) ([]domain_model.PhishingSimulationUserVulnerability, error)
	UpdateUserVulnerability(vulnPatch *domain_model.PhishingSimulationUserVulnerabilityPatch) error
	GetRecognitionFeatures() ([]domain_model.PhishingSimulationRecognitionFeature, error)
}
