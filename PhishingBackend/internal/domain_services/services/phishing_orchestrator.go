package services

import (
	"log/slog"
	"math/rand"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"time"
)

var _ PhishingOrchestrator = (*PhishingOrchestratorImpl)(nil)

type PhishingOrchestrator interface {
	StartPhishingRunGenerationJob()
	StartPhishingRunStregthDetectionJob()
}

type PhishingOrchestratorImpl struct {
	UserRepository               repositories.UserRepository
	PhishingSimulationRepository repositories.PhishingSimulationRepository
	PhishingRunService           PhishingRunService
}

func (p *PhishingOrchestratorImpl) StartPhishingRunGenerationJob() {
	go StartRandomCronJob(2*time.Hour, 3*time.Hour, p.generatePhishingRuns)
}

func (p *PhishingOrchestratorImpl) StartPhishingRunStregthDetectionJob() {
	go StartCronStyleJob("* */6 * * *", p.detectUserStrengths)
}

func (p *PhishingOrchestratorImpl) generatePhishingRuns(currentTime time.Time) {
	users, err := p.UserRepository.GetUsersForPhishingSimulation()
	if err != nil {
		return
	}

	day := 24 * time.Hour

	minPeriod := 5 * day
	maxPeriod := 14 * day
	for _, user := range users {
		periodUntilNextRun := minPeriod + time.Duration(rand.Int63n(int64(maxPeriod-minPeriod)))

		latestRun, err := p.PhishingSimulationRepository.GetLatestRun(user.ID)
		if err != nil {
			continue
		}

		if latestRun == nil ||
			latestRun.Email == nil ||
			latestRun.Email.SentAt == nil ||
			latestRun.Email.SentAt.Add(periodUntilNextRun).UTC().Before(currentTime) {
			err = p.PhishingRunService.GenerateRun(&user)
			if err != nil {
				slog.Error("Generate Run Failed", "error", err)
			}
		}
	}
}

func (p *PhishingOrchestratorImpl) detectUserStrengths(currentTime time.Time) {
	day := 24 * time.Hour

	unprocessedRuns, _ := p.PhishingSimulationRepository.GetUnprocessedRuns()
	toBeProcessedAfter := 4 * day
	for _, run := range unprocessedRuns {
		if run.ProcessedAt == nil &&
			run.Email != nil &&
			run.Email.SentAt != nil &&
			run.Email.SentAt.Add(toBeProcessedAfter).UTC().Before(currentTime) {
			slog.Info("Process Unclicked Run", "info", run)
			err := p.PhishingRunService.ProcessUnclickedRun(&run)
			if err != nil {
				slog.Error("Processing Unclicked Run Failed", "error", err)
			}
		}
	}
}
