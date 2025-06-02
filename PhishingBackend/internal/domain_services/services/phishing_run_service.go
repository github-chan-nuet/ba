package services

import (
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/email"
)

var _ PhishingRunService = (*PhishingRunServiceImpl)(nil)

type PhishingRunService interface {
	GenerateRun(*domain_model.User) error
}

type PhishingRunServiceImpl struct {
	EmailSender email.EmailSender
}

func (s *PhishingRunServiceImpl) GenerateRun(*domain_model.User) error {

	return nil
}
