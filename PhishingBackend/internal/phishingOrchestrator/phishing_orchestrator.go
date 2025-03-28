package orchestrator

import (
	cronjob "phishing_backend/internal/cronJob"
	"time"
)

func StartPhishingOrchestrationJob() {
	go cronjob.Start(time.Hour, orchestratePhishingSimulation)
}

func orchestratePhishingSimulation(_ time.Time) {

}
