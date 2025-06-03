package services

import (
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/email"
	"phishing_backend/internal/domain_services/interfaces/repositories"

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
		}
		err = s.PhishingSimulationRepository.Create(&run)
		if err != nil {
			return err
		}

		email := domain_model.Email{
			Content:   first.Content,
			Recipient: user.Email,
			Subject:   first.Subject,
		}
		err = s.EmailSender.Send(&email)
		if err != nil {
			return err
		}
	}
	return nil
}
