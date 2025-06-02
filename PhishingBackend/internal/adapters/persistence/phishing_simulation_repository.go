package persistence

import (
	"errors"
	"log/slog"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ repositories.PhishingSimulationRepository = (*PhishingSimulationRepositoryImpl)(nil)

type PhishingSimulationRepositoryImpl struct {
}

func (r *PhishingSimulationRepositoryImpl) Create(run *domain_model.PhishingSimulationRun) error {
	result := db.Create(run)
	if result.Error != nil {
		slog.Error("Could not save phishing simulation run")
	}
	return result.Error
}

func (r *PhishingSimulationRepositoryImpl) Update(run *domain_model.PhishingSimulationRun) error {
	return nil
}

func (r *PhishingSimulationRepositoryImpl) GetLatestRun(userId uuid.UUID) (*domain_model.PhishingSimulationRun, error) {
	var latestRun domain_model.PhishingSimulationRun
	result := db.Model(&domain_model.PhishingSimulationRun{}).
		Where("user_fk = ?", userId).
		First(&latestRun)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		slog.Error("Could not get latest run by user", "err", result.Error)
		return nil, result.Error
	}
	return &latestRun, nil
}

func (r *PhishingSimulationRepositoryImpl) GetTemplates() ([]domain_model.PhishingSimulationContentTemplate, error) {
	var templates []domain_model.PhishingSimulationContentTemplate
	result := db.Model(&domain_model.PhishingSimulationContentTemplate{}).Find(&templates)
	if result.Error != nil {
		slog.Error("Could not fetch phishing simulation templates")
		return nil, result.Error
	}
	return templates, nil
}

func (r *PhishingSimulationRepositoryImpl) GetUserVulnerabilities(userId uuid.UUID) ([]domain_model.PhishingSimulationUserVulnerability, error) {
	return nil, nil
}
