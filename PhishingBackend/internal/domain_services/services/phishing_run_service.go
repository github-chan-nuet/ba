package services

import (
	"errors"
	"log/slog"
	"math"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/email"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"time"

	"github.com/google/uuid"
)

var _ PhishingRunService = (*PhishingRunServiceImpl)(nil)

type PhishingRunService interface {
	GenerateRun(*domain_model.User) error
	TrackRunClick(*domain_model.PhishingSimulationRun) error
}

type PhishingRunServiceImpl struct {
	EmailRepository                repositories.EmailRepository
	EmailSender                    email.EmailSender
	PhishingSimulationRepository   repositories.PhishingSimulationRepository
	PhishingEmailGenerationService PhishingEmailGenerationService
}

func (s *PhishingRunServiceImpl) GenerateRun(user *domain_model.User) error {
	vulnerabilities, err := s.PhishingSimulationRepository.GetUserVulnerabilities(user.ID)
	if err != nil {
		return errors.New("Error while fetching vulnerabilities")
	}

	scoredTemplates := s.getScoredTemplates(vulnerabilities)
	if len(scoredTemplates) == 0 {
		return errors.New("No Scored Templates found")
	}

	lowestScoredTemplate := scoredTemplates[0]
	for _, scoredTemplate := range scoredTemplates[1:] {
		if scoredTemplate.Score < lowestScoredTemplate.Score {
			lowestScoredTemplate = scoredTemplate
		}
	}

	run := domain_model.PhishingSimulationRun{
		ID:                       uuid.New(),
		UserFk:                   user.ID,
		User:                     user,
		TemplateFk:               lowestScoredTemplate.Template.ID,
		Template:                 lowestScoredTemplate.Template,
		RecognitionFeatureValues: lowestScoredTemplate.RecognitionFeatureValues,
	}

	err = s.PhishingSimulationRepository.Create(&run)
	if err != nil {
		return err
	}

	email, err := s.sendRun(&run)
	if err != nil {
		return err
	}

	runPatch := domain_model.PhishingSimulationRunPatch{
		ID:      run.ID,
		EmailFk: &email.ID,
	}
	return s.PhishingSimulationRepository.Update(&runPatch)
}

func (s *PhishingRunServiceImpl) getScoredTemplates(vulnerabilities []domain_model.PhishingSimulationUserVulnerability) []domain_model.ScoredTemplate {
	var scoredTemplatesTotal []domain_model.ScoredTemplate
	recognitionFeatures, err := s.PhishingSimulationRepository.GetRecognitionFeatures()
	if err == nil {
		templates, err := s.PhishingSimulationRepository.GetTemplates()
		if err == nil {
			for _, template := range templates {
				scoredTemplates := template.GetScoredCombinations(vulnerabilities, recognitionFeatures)
				scoredTemplatesTotal = append(scoredTemplatesTotal, scoredTemplates...)
			}
		}
	}
	return scoredTemplatesTotal
}

func (s *PhishingRunServiceImpl) sendRun(run *domain_model.PhishingSimulationRun) (*domain_model.Email, error) {
	email := s.PhishingEmailGenerationService.GenerateEmail(run)
	if email == nil {
		return email, errors.New("Email could not be generated")
	}

	err := s.EmailRepository.Create(email)
	if err != nil {
		return email, err
	}

	err = s.EmailSender.Send(email)
	if err != nil {
		return email, err
	}

	now := time.Now().UTC()
	emailPatch := domain_model.EmailPatch{
		ID:     email.ID,
		SentAt: &now,
	}
	email.SentAt = &now

	err = s.EmailRepository.Update(&emailPatch)
	if err != nil {
		return email, err
	}

	return email, nil
}

func (s *PhishingRunServiceImpl) TrackRunClick(run *domain_model.PhishingSimulationRun) error {
	if run.Email.ClickedAt != nil {
		return errors.New("Run Click already tracked")
	}

	now := time.Now().UTC()
	emailPatch := domain_model.EmailPatch{
		ID:        run.Email.ID,
		ClickedAt: &now,
	}
	err := s.EmailRepository.Update(&emailPatch)
	if err != nil {
		return err
	}

	vulnerabilities, err := s.PhishingSimulationRepository.GetUserVulnerabilities(run.User.ID)
	if err != nil {
		return errors.New("Error while fetching vulnerabilities")
	}
	for _, featVal := range run.RecognitionFeatureValues {
		found := false
		for _, vuln := range vulnerabilities {
			if vuln.ContentCategory.ID == run.Template.ContentCategory.ID && vuln.RecognitionFeature.ID == featVal.RecognitionFeature.ID {
				vulnPatch := &domain_model.PhishingSimulationUserVulnerabilityPatch{
					ID:    vuln.ID,
					Score: float32(math.Max(float64(vuln.Score)-0.5, 1)),
				}
				err := s.PhishingSimulationRepository.UpdateUserVulnerability(vulnPatch)
				if err != nil {
					slog.Error("Tracking Run Click for User Vulnerability with ID " + vuln.ID.String() + " failed")
				}
				found = true
				break
			}
		}
		if !found {
			vuln := domain_model.PhishingSimulationUserVulnerability{
				ID:                   uuid.New(),
				Score:                1,
				UserFk:               run.UserFk,
				ContentCategoryFk:    run.Template.ContentCategoryFk,
				RecognitionFeatureFk: featVal.RecognitionFeatureFk,
			}
			err := s.PhishingSimulationRepository.CreateUserVulnerability(&vuln)
			if err != nil {
				slog.Error("Tracking Run Click for New User Vulnerability failed")
			}
		}
	}
	return nil
}
