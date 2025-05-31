package services

import (
	"math/rand"
	"phishing_backend/internal/domain_services/interfaces/email"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"time"
)

var _ PhishingOrchestrator = (*PhishingOrchestratorImpl)(nil)

type PhishingOrchestrator interface {
	StartPhishingJob(days time.Duration)
}

type PhishingOrchestratorImpl struct {
	EmailSender                  email.EmailSender
	PhishingSimulationRepository repositories.PhishingSimulationRepository
}

func (r *PhishingOrchestratorImpl) StartPhishingJob(days time.Duration) {
	//d := r.CalculateDuration(days)
	//r.
}

func (r *PhishingOrchestratorImpl) CalculateDuration(days int) time.Duration {
	randMinutes := rand.Intn(60*24 - 1)
	duration := time.Duration(int64(days-1))*time.Hour*24 + time.Duration(randMinutes)*time.Minute
	return duration
}
