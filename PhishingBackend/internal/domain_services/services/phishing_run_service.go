package services

import "phishing_backend/internal/domain_model"

var _ PhishingRunService = (*PhishingRunServiceImpl)(nil)

type PhishingRunService interface {
	GenerateRun(*domain_model.User) error
}

type PhishingRunServiceImpl struct {
}

func (s *PhishingRunServiceImpl) GenerateRun(*domain_model.User) error {
	return nil
}
