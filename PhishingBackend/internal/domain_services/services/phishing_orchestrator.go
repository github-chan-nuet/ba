package services

import (
	"math/rand"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"time"
)

var _ PhishingOrchestrator = (*PhishingOrchestratorImpl)(nil)

type PhishingOrchestrator interface {
	StartPhishingJob()
}

type PhishingOrchestratorImpl struct {
	UserRepository               repositories.UserRepository
	PhishingSimulationRepository repositories.PhishingSimulationRepository
	PhishingRunService           PhishingRunService
}

func (p *PhishingOrchestratorImpl) StartPhishingJob() {
	go StartRandomCronJob(15*time.Minute, 60*time.Minute, p.generatePhishingRuns)
}

func (p *PhishingOrchestratorImpl) generatePhishingRuns(currentTime time.Time) {
	users, err := p.UserRepository.GetUsersForPhishingSimulation()
	if err != nil {
		return
	}

	day := 24 * time.Hour

	minPeriod := 2 * day
	maxPeriod := 10 * day
	for _, user := range users {
		periodUntilNextRun := minPeriod + time.Duration(rand.Int63n(int64(maxPeriod-minPeriod)))

		latestRun, err := p.PhishingSimulationRepository.GetLatestRun(user.ID)
		if err != nil {
			continue
		}

		if latestRun == nil ||
			latestRun.SentAt == nil ||
			latestRun.SentAt.Add(periodUntilNextRun).UTC().Before(currentTime) {
			err = p.PhishingRunService.GenerateRun(&user)
		}
	}
}
