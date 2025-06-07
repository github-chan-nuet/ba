package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/email"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"strconv"
	"strings"
	"testing"
	"text/template"
	"time"
)

type EmailSenderMock struct {
	sentEmail *domain_model.Email
}

func (e *EmailSenderMock) Send(email *domain_model.Email) error {
	e.sentEmail = email
	return nil
}

var _ email.EmailSender = (*EmailSenderMock)(nil)

type reminderOrchestratorContext struct {
	sut          ReminderOrchestratorImpl
	lessonRepo   *repositories.MockLessonCompletionRepository
	reminderRepo *repositories.MockReminderRepository
	usrRepo      *repositories.MockUserRepository
	templateRepo *repositories.MockReminderEmailTemplateRepository
	emailSender  *EmailSenderMock
}

func createReminderOrchestratorSut(t *testing.T) *reminderOrchestratorContext {
	ctrl := gomock.NewController(t)
	ctx := reminderOrchestratorContext{
		lessonRepo:   repositories.NewMockLessonCompletionRepository(ctrl),
		reminderRepo: repositories.NewMockReminderRepository(ctrl),
		usrRepo:      repositories.NewMockUserRepository(ctrl),
		templateRepo: repositories.NewMockReminderEmailTemplateRepository(ctrl),
		emailSender:  &EmailSenderMock{},
	}
	ctx.sut = ReminderOrchestratorImpl{
		StartEachDayAfter:          0,
		EmailSender:                ctx.emailSender,
		LessonCompletionRepository: ctx.lessonRepo,
		ReminderRepository:         ctx.reminderRepo,
		UserRepository:             ctx.usrRepo,
		TemplateRepository:         ctx.templateRepo,
	}
	return &ctx
}

func mockTemplate(repo *repositories.MockReminderEmailTemplateRepository) *domain_model.ReminderEmailTemplate {
	templ := domain_model.ReminderEmailTemplate{
		Id:       2,
		Template: "Hallo {{ .Firstname }} {{  .Lastname }}",
		Subject:  "Testemail",
	}
	repo.EXPECT().GetAll().Return(&[]domain_model.ReminderEmailTemplate{templ}, nil)
	return &templ
}

func mockTemplates(repo *repositories.MockReminderEmailTemplateRepository) *domain_model.ReminderEmailTemplate {
	first := domain_model.ReminderEmailTemplate{
		Id:       1,
		Template: "Error",
		Subject:  "Error",
	}
	templ := domain_model.ReminderEmailTemplate{
		Id:       2,
		Template: "Hallo {{ .Firstname }} {{  .Lastname }}",
		Subject:  "Testemail",
	}
	repo.EXPECT().GetAll().Return(&[]domain_model.ReminderEmailTemplate{first, templ}, nil)
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

func TestShouldSendReminderWhenNoReminderWasSentAndNoLessonCompleted(t *testing.T) {
	// given
	ctx := createReminderOrchestratorSut(t)
	user := createUser()
	compLessons := make(map[uuid.UUID]time.Time)
	ctx.lessonRepo.EXPECT().GetLatestLessonCompletions().Return(compLessons, nil)
	ctx.usrRepo.EXPECT().GetAllUsers().Return(&[]domain_model.User{user}, nil)
	ctx.reminderRepo.EXPECT().GetAll().Return([]domain_model.Reminder{}, nil)
	templ := mockTemplate(ctx.templateRepo)
	ctx.reminderRepo.EXPECT().SaveOrUpdate(gomock.Any()).Return(nil)

	// when
	ctx.sut.ExecuteReminderJob()

	// then
	sentMail := ctx.emailSender.sentEmail
	wantMail := getWantMail(user, templ)
	assert.Equal(t, wantMail, *sentMail)
}

func TestShouldSendReminderWhenLessonWasCompletedOneWeekAgo(t *testing.T) {
	// given
	ctx := createReminderOrchestratorSut(t)
	user := createUser()
	compLessons := make(map[uuid.UUID]time.Time)
	compLessons[user.ID] = time.Now().UTC().Add(-7*24*time.Hour - 5*time.Minute)
	ctx.lessonRepo.EXPECT().GetLatestLessonCompletions().Return(compLessons, nil)
	ctx.usrRepo.EXPECT().GetAllUsers().Return(&[]domain_model.User{user}, nil)
	ctx.reminderRepo.EXPECT().GetAll().Return([]domain_model.Reminder{}, nil)
	templ := mockTemplate(ctx.templateRepo)
	ctx.reminderRepo.EXPECT().SaveOrUpdate(gomock.Any()).Return(nil)

	// when
	ctx.sut.ExecuteReminderJob()

	// then
	sentMail := ctx.emailSender.sentEmail
	wantMail := getWantMail(user, templ)
	assert.Equal(t, wantMail, *sentMail)
}

func TestShouldNotSendReminderWhen2RemindersWereSent(t *testing.T) {
	// given
	ctx := createReminderOrchestratorSut(t)
	user := createUser()
	compLessons := make(map[uuid.UUID]time.Time)
	ctx.lessonRepo.EXPECT().GetLatestLessonCompletions().Return(compLessons, nil)
	ctx.usrRepo.EXPECT().GetAllUsers().Return(&[]domain_model.User{user}, nil)
	reminder := domain_model.Reminder{
		ID:         uuid.New(),
		UserFk:     user.ID,
		SentTime:   time.Now().UTC().Add(-2 * 7 * 24 * time.Hour),
		Count:      2,
		TemplateFk: 1,
	}
	ctx.reminderRepo.EXPECT().GetAll().Return([]domain_model.Reminder{reminder}, nil)
	mockTemplate(ctx.templateRepo)

	// when
	ctx.sut.ExecuteReminderJob()

	// then
	sentMail := ctx.emailSender.sentEmail
	assert.Nil(t, sentMail)
}

func TestShouldNotSendReminderWhenLessonWasCompletedAlmostOneWeekAgo(t *testing.T) {
	// given
	ctx := createReminderOrchestratorSut(t)
	user := createUser()
	compLessons := make(map[uuid.UUID]time.Time)
	compLessons[user.ID] = time.Now().UTC().Add(-7*24*time.Hour + 5*time.Minute)
	ctx.lessonRepo.EXPECT().GetLatestLessonCompletions().Return(compLessons, nil)
	ctx.usrRepo.EXPECT().GetAllUsers().Return(&[]domain_model.User{user}, nil)
	ctx.reminderRepo.EXPECT().GetAll().Return([]domain_model.Reminder{}, nil)
	mockTemplate(ctx.templateRepo)

	// when
	ctx.sut.ExecuteReminderJob()

	// then
	sentMail := ctx.emailSender.sentEmail
	assert.Nil(t, sentMail)
}

func TestShouldNotSendReminderWhenLessonWasCompletedTwoWeeksAgoAndTwoRemindersWereSent(t *testing.T) {
	// given
	ctx := createReminderOrchestratorSut(t)
	user := createUser()
	compLessons := make(map[uuid.UUID]time.Time)
	compLessons[user.ID] = time.Now().UTC().Add(-2 * 7 * 24 * time.Hour)
	ctx.lessonRepo.EXPECT().GetLatestLessonCompletions().Return(compLessons, nil)
	ctx.usrRepo.EXPECT().GetAllUsers().Return(&[]domain_model.User{user}, nil)
	reminder := domain_model.Reminder{
		ID:         uuid.New(),
		UserFk:     user.ID,
		SentTime:   time.Now().UTC().Add(-2 * 7 * 24 * time.Hour),
		Count:      2,
		TemplateFk: 1,
	}
	ctx.reminderRepo.EXPECT().GetAll().Return([]domain_model.Reminder{reminder}, nil)
	mockTemplate(ctx.templateRepo)

	// when
	ctx.sut.ExecuteReminderJob()

	// then
	sentMail := ctx.emailSender.sentEmail
	assert.Nil(t, sentMail)
}

func TestShouldSendEmailWithNewTemplate(t *testing.T) {
	// given
	ctx := createReminderOrchestratorSut(t)
	user := createUser()
	compLessons := make(map[uuid.UUID]time.Time)
	compLessons[user.ID] = time.Now().UTC().Add(-2 * 7 * 24 * time.Hour)
	ctx.lessonRepo.EXPECT().GetLatestLessonCompletions().Return(compLessons, nil)
	ctx.usrRepo.EXPECT().GetAllUsers().Return(&[]domain_model.User{user}, nil)
	reminder := domain_model.Reminder{
		ID:         uuid.New(),
		UserFk:     user.ID,
		SentTime:   time.Now().UTC().Add(-2 * 7 * 24 * time.Hour),
		Count:      1,
		TemplateFk: 1,
	}
	ctx.reminderRepo.EXPECT().GetAll().Return([]domain_model.Reminder{reminder}, nil)
	// mocks two templates
	templ := mockTemplates(ctx.templateRepo)
	ctx.reminderRepo.EXPECT().SaveOrUpdate(gomock.Any()).Return(nil)

	// when
	ctx.sut.ExecuteReminderJob()

	// then
	sentMail := ctx.emailSender.sentEmail
	wantMail := getWantMail(user, templ)
	assert.Equal(t, wantMail, *sentMail)
}

func TestShouldNotSendEmailWhenReminderWasSentAlmostOneWeekAgo(t *testing.T) {
	// given
	ctx := createReminderOrchestratorSut(t)
	user := createUser()
	compLessons := make(map[uuid.UUID]time.Time)
	compLessons[user.ID] = time.Now().UTC().Add(-2 * 7 * 24 * time.Hour)
	ctx.lessonRepo.EXPECT().GetLatestLessonCompletions().Return(compLessons, nil)
	ctx.usrRepo.EXPECT().GetAllUsers().Return(&[]domain_model.User{user}, nil)
	reminder := domain_model.Reminder{
		ID:         uuid.New(),
		UserFk:     user.ID,
		SentTime:   time.Now().UTC().Add(-7*24*time.Hour + 5*time.Minute),
		Count:      1,
		TemplateFk: 1,
	}
	ctx.reminderRepo.EXPECT().GetAll().Return([]domain_model.Reminder{reminder}, nil)
	mockTemplate(ctx.templateRepo)

	// when
	ctx.sut.ExecuteReminderJob()

	// then
	sentMail := ctx.emailSender.sentEmail
	assert.Nil(t, sentMail)
}

func createUser() domain_model.User {
	return domain_model.User{
		ID:        uuid.UUID{},
		Firstname: "Hans",
		Lastname:  "Nötig",
		Email:     "a@gmail.com",
	}
}
