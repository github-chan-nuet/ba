package services

import (
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
	EmailSender                  email.EmailSender
	PhishingSimulationRepository repositories.PhishingSimulationRepository
}

func (s *PhishingRunServiceImpl) GenerateRun(user *domain_model.User) error {
	scoredTemplates := s.getScoredTemplates(user.ID)
	slog.Info("scored Templates", "info", scoredTemplates)

	templates, err := s.PhishingSimulationRepository.GetTemplates()
	if err != nil {
		return err
	}
	if len(templates) > 0 {
		first := templates[0]

		run := domain_model.PhishingSimulationRun{
			ID:         uuid.New(),
			UserFk:     user.ID,
			TemplateFk: first.ID,
			Subject:    first.Subject,
			Content:    first.Content,
		}

		err = s.PhishingSimulationRepository.Create(&run)
		if err != nil {
			return err
		}

		return nil //s.sendRun(&run)
	}
	return nil
}

func (s *PhishingRunServiceImpl) getScoredTemplates(userId uuid.UUID) []domain_model.ScoredTemplate {
	var scoredTemplatesTotal []domain_model.ScoredTemplate
	vulnerabilities, err := s.PhishingSimulationRepository.GetUserVulnerabilities(userId)
	slog.Info("vulnerabilities", "info", vulnerabilities)
	if err == nil {
		recognitionFeatures, err := s.PhishingSimulationRepository.GetRecognitionFeatures()
		if err == nil {
			templates, err := s.PhishingSimulationRepository.GetTemplates()
			slog.Info("templates", "info", templates)
			if err == nil {
				for _, template := range templates {
					scoredTemplates := template.GetScoredCombinations(vulnerabilities, recognitionFeatures)
					scoredTemplatesTotal = append(scoredTemplatesTotal, scoredTemplates...)
				}
			}
		}
	}
	return scoredTemplatesTotal
}

func (s *PhishingRunServiceImpl) sendRun(run *domain_model.PhishingSimulationRun) error {
	email := domain_model.Email{
		Subject:   run.Subject,
		Content:   run.Content,
		Recipient: run.User.Email,
	}

	err := s.EmailSender.Send(&email)
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
