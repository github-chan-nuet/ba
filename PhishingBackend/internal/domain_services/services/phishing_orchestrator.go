package services

import (
	"time"
)

func StartPhishingOrchestrationJob() {
	go StartCronJob(time.Hour, orchestratePhishingSimulation)
}

func orchestratePhishingSimulation(_ time.Time) {

}
