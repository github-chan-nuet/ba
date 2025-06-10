package adapters

import (
	"log/slog"
	"net/smtp"
	"os"
	"phishing_backend/internal/adapters/communication"
	"phishing_backend/internal/adapters/persistence"
	"phishing_backend/internal/domain_services/interfaces/email"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"phishing_backend/internal/domain_services/services"
	"strconv"
	"time"
)

type Dependencies struct {
	// repositories
	UserRepository               repositories.UserRepository
	LessonCompletionRepository   repositories.LessonCompletionRepository
	ExamRepository               repositories.ExamRepository
	ExamCompletionRepository     repositories.ExamCompletionRepository
	ReminderRepository           repositories.ReminderRepository
	ReminderTemplateRepository   repositories.ReminderEmailTemplateRepository
	EmailRepository              repositories.EmailRepository
	PhishingSimulationRepository repositories.PhishingSimulationRepository

	// services
	EmailSender                    email.EmailSender
	ExperienceService              services.ExperienceService
	ExamCompletionService          services.ExamCompletionService
	Authenticator                  services.Authenticator
	ReminderOrchestrator           services.ReminderOrchestrator
	LessonCompletionService        services.LessonCompletionService
	UserService                    services.UserService
	PhishingEmailGenerationService services.PhishingEmailGenerationService
	PhishingRunService             services.PhishingRunService
	PhishingOrchestrator           services.PhishingOrchestrator
}

func ResolveDependencies() *Dependencies {
	d := &Dependencies{}
	// repositories
	d.UserRepository = &persistence.UserRepositoryImpl{}
	d.LessonCompletionRepository = &persistence.LessonCompletionRepositoryImpl{}
	d.ExamRepository = &persistence.ExamRepositoryImpl{}
	d.ExamCompletionRepository = &persistence.ExamCompletionRepositoryImpl{}
	d.ReminderRepository = &persistence.ReminderRepositoryImpl{}
	d.ReminderTemplateRepository = &persistence.ReminderEmailTemplateRepositoryImpl{}
	d.EmailRepository = &persistence.EmailRepositoryImpl{}
	d.PhishingSimulationRepository = &persistence.PhishingSimulationRepositoryImpl{}

	// services
	d.EmailSender = &communication.EmailSenderImpl{
		SmtpUser:   os.Getenv("PHBA_SMTP_USER"),
		SmtpPw:     os.Getenv("PHBA_SMTP_PASSWORD"),
		SmtpAddr:   os.Getenv("PHBA_SMTP_ADDR"),
		SmtpHost:   os.Getenv("PHBA_SMTP_HOST"),
		SendMailFn: smtp.SendMail,
	}
	d.ExperienceService = &services.ExperienceServiceImpl{
		LessonCompRepo: d.LessonCompletionRepository,
		ExamCompRepo:   d.ExamCompletionRepository,
	}
	d.ExamCompletionService = &services.ExamCompletionServiceImpl{
		ExamRepo:          d.ExamRepository,
		ExamCompRepo:      d.ExamCompletionRepository,
		ExperienceService: d.ExperienceService,
	}
	d.Authenticator = &services.AuthenticatorImpl{
		UserRepository: d.UserRepository,
	}
	d.LessonCompletionService = &services.LessonCompletionServiceImpl{
		Repo: d.LessonCompletionRepository,
	}
	d.UserService = &services.UserServiceImpl{
		UserRepository: d.UserRepository,
	}
	d.PhishingEmailGenerationService = &services.PhishingEmailGenerationServiceImpl{}
	d.PhishingRunService = &services.PhishingRunServiceImpl{
		EmailRepository:                d.EmailRepository,
		EmailSender:                    d.EmailSender,
		PhishingSimulationRepository:   d.PhishingSimulationRepository,
		PhishingEmailGenerationService: d.PhishingEmailGenerationService,
	}
	setReminderOrchestrator(d)
	setPhishingOrchestrator(d)
	return d
}

func setReminderOrchestrator(d *Dependencies) {
	envName := "PHBA_REMINDER_DELAY"
	reminderDelayStr := os.Getenv(envName)
	i, err := strconv.ParseInt(reminderDelayStr, 10, 64)
	if err != nil {
		slog.Error(envName+" must be an integer but is not", envName, reminderDelayStr, "error", err)
		panic(envName + " must be an integer but is not")
	}
	d.ReminderOrchestrator = &services.ReminderOrchestratorImpl{
		StartEachDayAfter:          time.Duration(i),
		EmailSender:                d.EmailSender,
		LessonCompletionRepository: d.LessonCompletionRepository,
		ReminderRepository:         d.ReminderRepository,
		UserRepository:             d.UserRepository,
		TemplateRepository:         d.ReminderTemplateRepository,
	}
}

func setPhishingOrchestrator(d *Dependencies) {
	d.PhishingOrchestrator = &services.PhishingOrchestratorImpl{
		UserRepository:               d.UserRepository,
		PhishingSimulationRepository: d.PhishingSimulationRepository,
		PhishingRunService:           d.PhishingRunService,
	}
}
