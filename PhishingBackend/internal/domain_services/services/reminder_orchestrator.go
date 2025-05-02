package services

import (
	"time"
)

var _ ReminderOrchestrator = (*ReminderOrchestratorImpl)(nil)

type ReminderOrchestrator interface {
	StartReminderJob(d time.Duration)
}

type ReminderOrchestratorImpl struct {
}

func (r *ReminderOrchestratorImpl) StartReminderJob(d time.Duration) {
	go StartCronJob(d, r.sendReminder)
}

func (r *ReminderOrchestratorImpl) sendReminder(utc time.Time) {

}
