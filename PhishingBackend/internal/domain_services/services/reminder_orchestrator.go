package services

import (
	"github.com/google/uuid"
	"log/slog"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/email"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"slices"
	"sync"
	"time"
)

var _ ReminderOrchestrator = (*ReminderOrchestratorImpl)(nil)

type ReminderOrchestrator interface {
	ExecuteReminderJobAfterDurationEachDay(d time.Duration)
}

type ReminderOrchestratorImpl struct {
	EmailSender                email.EmailSender
	LessonCompletionRepository repositories.LessonCompletionRepository
	ReminderRepository         repositories.ReminderRepository
	UserRepository             repositories.UserRepository
}

func (r *ReminderOrchestratorImpl) ExecuteReminderJobAfterDurationEachDay(d time.Duration) {
	for {
		nextInvocation := time.Now().Truncate(time.Hour * 24).Add(time.Hour*24 + d)
		timeTillNext := nextInvocation.Sub(time.Now())
		var wg sync.WaitGroup
		wg.Add(1)
		time.AfterFunc(timeTillNext, func() {
			r.sendReminder()
			wg.Done()
		})
		wg.Wait()
	}
}

func todo() {
	//smtpUser := os.Getenv("PHBA_SMTP_USER")
	//smtpPw := os.Getenv("PHBA_SMTP_PASSWORD")
	//smtpAddr := os.Getenv("PHBA_SMTP_ADDR")
	//smtpHost := os.Getenv("PHBA_SMTP_HOST")
	//
	//mailer := communication.EmailSenderImpl{
	//	SmtpUser: smtpUser,
	//	SmtpPw:   smtpPw,
	//	SmtpAddr: smtpAddr,
	//	SmtpHost: smtpHost,
	//}
}

func sendMail() {
	//reminder := domain_model.Email{
	//	Content:   "Dies ist ein Test-Reminder",
	//	Recipient: "mischa.binder@stafag.ch",
	//	Subject:   "Test-Reminder | Securaware",
	//}
	//r.EmailSender.Send(&reminder)
}

func (r *ReminderOrchestratorImpl) sendReminder() {
	lasts, err := r.LessonCompletionRepository.GetLatestLessonCompletions()
	if err != nil {
		slog.Error("Skipping reminder job as lesson completions could not be get")
		return
	}
	userIds, err := r.UserRepository.GetAllUserIds()
	if err != nil {
		slog.Error("Skipping reminder job as user ids could not be retrieved")
		return
	}
	reminders, err := r.ReminderRepository.GetAll()
	if err != nil {
		slog.Error("Skipping reminder job as reminders")
		return
	}
	toExclude := r.getUsersThatHaveReceived2Reminders(&reminders)
	remindersToSend := make([]userReminderToSend, 0)
	for _, userId := range *userIds {
		slices.Contains(toExclude, userId)
	}
}

type userReminderToSend struct {
	user     domain_model.User
	reminder *domain_model.Reminder
}

func (r *ReminderOrchestratorImpl) getUsersThatHaveReceived2Reminders(reminders *[]domain_model.Reminder) *[]uuid.UUID {
	toExclude := make([]uuid.UUID, 0)
	for _, rem := range *reminders {
		if rem.Count >= 2 {
			toExclude = append(toExclude, rem.UserFk)
		}
	}
	return &toExclude
}
