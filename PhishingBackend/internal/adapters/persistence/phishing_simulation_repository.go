package persistence

import (
	"errors"
	"log/slog"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

var _ repositories.PhishingSimulationRepository = (*PhishingSimulationRepositoryImpl)(nil)

const uniqueUserVulnerability = "unique_vulnerability_per_user"

type PhishingSimulationRepositoryImpl struct {
}

func (r *PhishingSimulationRepositoryImpl) Create(run *domain_model.PhishingSimulationRun) error {
	result := db.Create(run)
	if result.Error != nil {
		slog.Error("Could not save phishing simulation run")
	}
	return result.Error
}

func (r *PhishingSimulationRepositoryImpl) Update(runPatch *domain_model.PhishingSimulationRunPatch) error {
	var existing domain_model.PhishingSimulationRun
	if err := db.First(&existing, runPatch.ID).Error; err != nil {
		return errors.New("Simulation run not found")
	}

	updates := map[string]interface{}{}
	if runPatch.EmailFk != nil {
		if existing.EmailFk != nil {
			return errors.New("EmailFk is already set")
		}
		updates["email_fk"] = *runPatch.EmailFk
	}
	if runPatch.ProcessedAt != nil {
		if existing.ProcessedAt != nil {
			return errors.New("ProcessedAt is already set")
		}
		updates["processed_at"] = *runPatch.ProcessedAt
	}

	if len(updates) > 0 {
		if err := db.Model(&domain_model.PhishingSimulationRun{}).
			Where("id = ?", runPatch.ID).
			Updates(updates).Error; err != nil {
			slog.Error("Could not update Phishing Simulation run")
			return err
		}
	}
	return nil
}

func (r *PhishingSimulationRepositoryImpl) GetRun(runId uuid.UUID) (*domain_model.PhishingSimulationRun, error) {
	var run domain_model.PhishingSimulationRun
	result := db.Model(&domain_model.PhishingSimulationRun{}).
		Preload("User").
		Preload("Template").
		Preload("Template.ContentCategory").
		Preload("RecognitionFeatureValues").
		Preload("RecognitionFeatureValues.RecognitionFeature").
		Preload("Email").
		Where("id = ?", runId).
		First(&run)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		slog.Error("Could not get run by id", "err", result.Error)
		return nil, result.Error
	}
	return &run, nil
}

func (r *PhishingSimulationRepositoryImpl) GetLatestRun(userId uuid.UUID) (*domain_model.PhishingSimulationRun, error) {
	var latestRun domain_model.PhishingSimulationRun
	result := db.Model(&domain_model.PhishingSimulationRun{}).
		Preload("User").
		Preload("Template").
		Preload("Template.ContentCategory").
		Preload("RecognitionFeatureValues").
		Preload("RecognitionFeatureValues.RecognitionFeature").
		Preload("Email").
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

func (r *PhishingSimulationRepositoryImpl) GetUnprocessedRuns() ([]domain_model.PhishingSimulationRun, error) {
	var runs []domain_model.PhishingSimulationRun
	result := db.Model(&domain_model.PhishingSimulationRun{}).
		Where("processed_at IS NULL").
		Preload("User").
		Preload("Template").
		Preload("Template.ContentCategory").
		Preload("RecognitionFeatureValues").
		Preload("RecognitionFeatureValues.RecognitionFeature").
		Preload("Email").
		Find(&runs)
	if result.Error != nil {
		slog.Error("Could not fetch unprocessed phishing simulation runs")
		return nil, result.Error
	}
	return runs, nil
}

func (r *PhishingSimulationRepositoryImpl) GetTemplates() ([]domain_model.PhishingSimulationContentTemplate, error) {
	var templates []domain_model.PhishingSimulationContentTemplate
	result := db.Model(&domain_model.PhishingSimulationContentTemplate{}).
		Preload("ContentCategory").
		Find(&templates)
	if result.Error != nil {
		slog.Error("Could not fetch phishing simulation templates")
		return nil, result.Error
	}
	return templates, nil
}

func (r *PhishingSimulationRepositoryImpl) CreateUserVulnerability(vulnerability *domain_model.PhishingSimulationUserVulnerability) error {
	result := db.Create(vulnerability)
	if result.Error != nil {
		var e *pgconn.PgError
		if errors.As(result.Error, &e) {
			if e.Code == "23505" && e.ConstraintName == uniqueUserVulnerability {
				return repositories.ErrUserVulnAlreadyExists
			}
		}
		slog.Error("Could not save user vulnerability", "err", result.Error)
	}
	return result.Error
}

func (r *PhishingSimulationRepositoryImpl) GetUserVulnerabilities(userId uuid.UUID) ([]domain_model.PhishingSimulationUserVulnerability, error) {
	var vulnerabilities []domain_model.PhishingSimulationUserVulnerability
	result := db.Model(&domain_model.PhishingSimulationUserVulnerability{}).
		Preload("ContentCategory").
		Preload("RecognitionFeature").
		Where("user_fk = ?", userId).
		Find(&vulnerabilities)
	if result.Error != nil {
		slog.Error("Could not fetch user vulnerabilities")
		return nil, result.Error
	}
	return vulnerabilities, nil
}

func (r *PhishingSimulationRepositoryImpl) UpdateUserVulnerability(vulnPatch *domain_model.PhishingSimulationUserVulnerabilityPatch) error {
	updates := make(map[string]interface{})
	updates["score"] = vulnPatch.Score

	result := db.Model(&domain_model.PhishingSimulationUserVulnerability{ID: vulnPatch.ID}).Updates(updates)
	if result.Error != nil {
		slog.Error("Could not update user vulnerability", "err", result.Error)
	}
	return result.Error
}

func (r *PhishingSimulationRepositoryImpl) GetRecognitionFeatures() ([]domain_model.PhishingSimulationRecognitionFeature, error) {
	var recognitionFeatures []domain_model.PhishingSimulationRecognitionFeature
	result := db.Model(&domain_model.PhishingSimulationRecognitionFeature{}).
		Preload("RecognitionFeatureValues").
		Preload("RecognitionFeatureValues.RecognitionFeature").
		Find(&recognitionFeatures)
	if result.Error != nil {
		slog.Error("Could not fetch phishing simulation recognition features")
		return nil, result.Error
	}
	return recognitionFeatures, nil
}
