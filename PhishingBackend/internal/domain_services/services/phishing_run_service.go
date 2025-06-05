package services

import (
	"errors"
	"log/slog"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/email"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"time"

	"github.com/google/uuid"
)

var _ PhishingRunService = (*PhishingRunServiceImpl)(nil)

type PhishingRunService interface {
	GenerateRun(*domain_model.User) error
}

type PhishingRunServiceImpl struct {
	EmailSender                    email.EmailSender
	PhishingSimulationRepository   repositories.PhishingSimulationRepository
	PhishingEmailGenerationService PhishingEmailGenerationService
}

func (s *PhishingRunServiceImpl) GenerateRun(user *domain_model.User) error {
	vulnerabilities, err := s.PhishingSimulationRepository.GetUserVulnerabilities(user.ID)
	if err != nil {
		return errors.New("Error while fetching vulnerabilities")
	}

	slog.Info("Vulnerabilities", "info", vulnerabilities)

	scoredTemplates := s.getScoredTemplates(vulnerabilities)
	if len(scoredTemplates) == 0 {
		return errors.New("No Scored Templates found")
	}

	slog.Info("Scored Templates", "info", scoredTemplates)

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

	return s.sendRun(&run)
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

func (s *PhishingRunServiceImpl) sendRun(run *domain_model.PhishingSimulationRun) error {
	email := s.PhishingEmailGenerationService.GenerateEmail(run)
	if email == nil {
		return errors.New("Email could not be generated")
	}

	err := s.EmailSender.Send(email)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	runPatch := domain_model.PhishingSimulationRunPatch{
		ID:     run.ID,
		SentAt: &now,
	}
	err = s.PhishingSimulationRepository.Update(&runPatch)
	if err != nil {
		return err
	}

	return nil
}
