//go:build integration

package reminder

import (
	random "crypto/rand"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"phishing_backend/integration_tests"
	"phishing_backend/internal/adapters/persistence"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/email"
	"phishing_backend/internal/domain_services/services"
	"strconv"
	"strings"
	"testing"
	"text/template"
)

func createRandomEmail() string {
	return strings.ToLower(random.Text()) + "@test.com"
}

var (
	userRepo       = persistence.UserRepositoryImpl{}
	templateRepo   = persistence.ReminderEmailTemplateRepositoryImpl{}
	reminderRepo   = persistence.ReminderRepositoryImpl{}
	lessonCompRepo = persistence.LessonCompletionRepositoryImpl{}
)

type EmailSenderMock struct {
	sentEmail []domain_model.Email
}

func (e *EmailSenderMock) Send(email *domain_model.Email) error {
	e.sentEmail = append(e.sentEmail, *email)
	return nil
}

var _ email.EmailSender = (*EmailSenderMock)(nil)

func createReminderOrchestrator() (*services.ReminderOrchestratorImpl, *EmailSenderMock) {
	emailSender := &EmailSenderMock{sentEmail: make([]domain_model.Email, 0, 10)}
	orch := services.ReminderOrchestratorImpl{
		StartEachDayAfter:          0,
		EmailSender:                emailSender,
		LessonCompletionRepository: &lessonCompRepo,
		ReminderRepository:         &reminderRepo,
		UserRepository:             &userRepo,
		TemplateRepository:         &templateRepo,
	}
	return &orch, emailSender
}

func createAndSaveTemplate() *domain_model.ReminderEmailTemplate {
	templ := domain_model.ReminderEmailTemplate{
		Id:       rand.Intn(1000),
		Template: "Hallo {{ .Firstname }} {{ .Lastname }}",
		Subject:  "Testemail",
	}
	integration_tests.GetDb().Create(&templ)
	return &templ
}

func getWantMail(u domain_model.User, t *domain_model.ReminderEmailTemplate) domain_model.Email {
	templ, _ := template.
		New(strconv.Itoa(t.Id)).
		Parse(t.Template)
	var sb strings.Builder
	templ.Execute(&sb, u)
	return domain_model.Email{
		Content:   sb.String(),
		Recipient: u.Email,
		Subject:   t.Subject,
	}
}

func TestShouldSendOutReminder(t *testing.T) {
	// given
	user := &domain_model.User{
		ID:                               uuid.New(),
		Firstname:                        "Hans",
		Lastname:                         "Nötig",
		Password:                         []byte("abcd"),
		Email:                            createRandomEmail(),
		ParticipatesInPhishingSimulation: false,
	}
	require.NoError(t, userRepo.CreateUser(user))
	templ := createAndSaveTemplate()

	sut, emailSender := createReminderOrchestrator()

	// when
	sut.ExecuteReminderJob()

	// then
	wantEmail := []domain_model.Email{getWantMail(*user, templ)}
	assert.ElementsMatch(t, wantEmail, emailSender.sentEmail)
}
