package services

import (
	"github.com/google/uuid"
	"log/slog"
	"math/rand"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/email"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"
)

var _ ReminderOrchestrator = (*ReminderOrchestratorImpl)(nil)

type ReminderOrchestrator interface {
	ExecuteReminderJobAfterDurationEachDay()
}

type ReminderOrchestratorImpl struct {
	StartEachDayAfter          time.Duration
	EmailSender                email.EmailSender
	LessonCompletionRepository repositories.LessonCompletionRepository
	ReminderRepository         repositories.ReminderRepository
	UserRepository             repositories.UserRepository
	TemplateRepository         repositories.ReminderEmailTemplateRepository
}

func (r *ReminderOrchestratorImpl) ExecuteReminderJobAfterDurationEachDay() {
	for {
		nextInvocation := time.Now().Truncate(time.Hour * 24).Add(time.Hour*24 + r.StartEachDayAfter)
		timeTillNext := nextInvocation.Sub(time.Now())
		var wg sync.WaitGroup
		wg.Add(1)
		time.AfterFunc(timeTillNext, func() {
			r.executeReminderJob()
			wg.Done()
		})
		wg.Wait()
	}
}

func (r *ReminderOrchestratorImpl) executeReminderJob() {
	lasts, err := r.LessonCompletionRepository.GetLatestLessonCompletions()
	if err != nil {
		slog.Error("Skipping reminder job as lesson completions could not be get")
		return
	}
	users, err := r.UserRepository.GetAllUsers()
	if err != nil {
		slog.Error("Skipping reminder job as user ids could not be retrieved")
		return
	}
	reminders, err := r.ReminderRepository.GetAll()
	userReminderMap := r.createUserReminderMap(&reminders)
	if err != nil {
		slog.Error("Skipping reminder job as reminders could not be fetched")
		return
	}
	toExclude := r.getUsersThatHaveReceived2Reminders(&reminders)
	toPrepares := make([]userReminderToPrepare, 0)
	now := time.Now().UTC()
	for _, user := range *users {
		if _, shouldExclude := toExclude[user.ID]; !shouldExclude {
			last, ok := lasts[user.ID]
			if !ok {
				toPrepares = append(toPrepares, userReminderToPrepare{user: user})
				continue
			}
			reminder := userReminderMap[user.ID]
			oneWeekAgo := now.Add(-1 * time.Hour * 24 * 7)
			if last.Before(oneWeekAgo) && reminder.SentTime.Before(oneWeekAgo) {
				toPrepare := userReminderToPrepare{user: user, reminder: userReminderMap[user.ID]}
				toPrepares = append(toPrepares, toPrepare)
			}
		}
	}
	r.prepareAndSendReminders(&toPrepares)
}

type userReminderToPrepare struct {
	user     domain_model.User
	reminder *domain_model.Reminder
}

func (r *ReminderOrchestratorImpl) prepareAndSendReminders(toPrepares *[]userReminderToPrepare) {
	templates, err := r.TemplateRepository.GetAll()
	if err != nil {
		slog.Error("Skipping reminder job as templates could not be retrieved")
		return
	}
	idToActualTemplate, err := r.prepareTemplates(templates)
	if err != nil {
		slog.Error("Skipping reminder job as a template could not be created")
		return
	}
	for _, toPrepare := range *toPrepares {
		var reminder domain_model.Reminder
		if toPrepare.reminder == nil {
			reminder = domain_model.Reminder{
				User:       &toPrepare.user,
				UserFk:     toPrepare.user.ID,
				SentTime:   time.Now().UTC(),
				Count:      1,
				TemplateFk: (*templates)[rand.Intn(len(*templates))].Id,
			}
		} else {
			reminder = *toPrepare.reminder
			reminder.SentTime = time.Now().UTC()
			reminder.Count = reminder.Count + 1
			reminder.User = &toPrepare.user
			newRandTemp := r.getNewRandomTemplate(templates, reminder.TemplateFk)
			reminder.EmailTemplate = newRandTemp
			reminder.TemplateFk = newRandTemp.Id
		}
		r.sendAndSaveReminder(&reminder, idToActualTemplate[reminder.TemplateFk])
	}
}

func (r *ReminderOrchestratorImpl) prepareTemplates(ts *[]domain_model.ReminderEmailTemplate) (map[int]*template.Template, error) {
	templates := make(map[int]*template.Template, len(*ts))
	for _, t := range *ts {
		tmpl, err := template.
			New(strconv.Itoa(t.Id)).
			Parse(t.Template)
		if err != nil {
			slog.Error("Template could not be parsed", "id", t.Id, "template", t.Template, "error", err)
			return nil, err
		}
		templates[t.Id] = tmpl
	}
	return templates, nil
}

func (r *ReminderOrchestratorImpl) getNewRandomTemplate(templates *[]domain_model.ReminderEmailTemplate, toExclude int) *domain_model.ReminderEmailTemplate {
	for {
		n := rand.Intn(len(*templates))
		t := (*templates)[n]
		if t.Id != toExclude {
			return &t
		}
	}
}

type userTemplate struct {
	firstname string
	lastname  string
}

func (r *ReminderOrchestratorImpl) sendAndSaveReminder(reminder *domain_model.Reminder, t *template.Template) {
	usrTmpl := userTemplate{
		firstname: reminder.User.Firstname,
		lastname:  reminder.User.Lastname,
	}
	var sb strings.Builder
	err := t.Execute(&sb, usrTmpl)
	if err != nil {
		slog.Error("Template could not be executed", "template name", t.Name(), "error", err)
		return
	}
	mail := domain_model.Email{
		Content:   sb.String(),
		Recipient: reminder.User.Email,
		Subject:   reminder.EmailTemplate.Subject,
	}
	err = r.EmailSender.Send(&mail)
	if err != nil {
		slog.Error("Reminder email could not be send", "error", err)
		return
	}
	err = r.ReminderRepository.SaveOrUpdate(reminder)
	if err != nil {
		slog.Error("Reminder could not be saved", "error", err)
	}
}

func (r *ReminderOrchestratorImpl) getUsersThatHaveReceived2Reminders(reminders *[]domain_model.Reminder) map[uuid.UUID]struct{} {
	toExclude := make(map[uuid.UUID]struct{})
	for _, rem := range *reminders {
		if rem.Count >= 2 {
			toExclude[rem.UserFk] = struct{}{}
		}
	}
	return toExclude
}

func (r *ReminderOrchestratorImpl) createUserReminderMap(reminders *[]domain_model.Reminder) map[uuid.UUID]*domain_model.Reminder {
	userReminderMap := make(map[uuid.UUID]*domain_model.Reminder, len(*reminders))
	for _, rem := range *reminders {
		userReminderMap[rem.UserFk] = &rem
	}
	return userReminderMap
}
