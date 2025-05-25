package presentation

import (
	"log/slog"
	"net/http"
	"os"
	"phishing_backend/internal/adapters/persistence"
	"phishing_backend/internal/adapters/presentation/controllers"
	"phishing_backend/internal/adapters/presentation/error_handling"
	"phishing_backend/internal/domain_services/services"
	"runtime"
)

func SetupHttpServer() {
	sMux := NewSecurawareServeMux()
	cors := CorsMiddleware{Handler: sMux}
	panicRec := PanicRecoveryMiddleware{Handler: &cors}
	addr := os.Getenv("PHBA_WEBSERVER_ADDR")
	slog.Info("Web server listening...", "address", addr)
	err := http.ListenAndServe(addr, &panicRec)
	slog.Error("Web server stopped", "error", err)
	os.Exit(1)
}

func NewSecurawareServeMux() *http.ServeMux {
	// repositories
	userRepository := persistence.UserRepositoryImpl{}
	lessonCompletionRepository := persistence.LessonCompletionRepositoryImpl{}
	examRepo := persistence.ExamRepositoryImpl{}
	examCompRepo := persistence.ExamCompletionRepositoryImpl{}

	// services
	expService := services.ExperienceServiceImpl{
		LessonCompRepo: &lessonCompletionRepository,
		ExamCompRepo:   &examCompRepo,
	}
	examCompService := services.ExamCompletionServiceImpl{
		ExamRepo:          &examRepo,
		ExamCompRepo:      &examCompRepo,
		ExperienceService: &expService,
	}
	authenticator := services.AuthenticatorImpl{
		UserRepository: &userRepository,
	}

	// controllers
	userController := controllers.UserController{
		Authenticator: &authenticator,
		UserService: &services.UserServiceImpl{
			UserRepository: &userRepository,
		},
		ExperienceService: &expService,
	}
	lessonCompletionController := controllers.LessonCompletionController{
		Authenticator: &authenticator,
		LessonCompletionService: &services.LessonCompletionServiceImpl{
			Repo: &lessonCompletionRepository,
		},
		ExperienceService:          &expService,
		LessonCompletionRepository: &lessonCompletionRepository,
	}
	examController := controllers.ExamController{
		Authenticator:         &authenticator,
		ExamRepository:        &examRepo,
		ExamCompRepo:          &examCompRepo,
		ExamCompletionService: &examCompService,
	}

	sMux := http.NewServeMux()

	// health
	sMux.HandleFunc("GET /api/health", controllers.GetHealth)

	// lesson completions
	sMux.HandleFunc("GET /api/courses/{courseId}/completions", lessonCompletionController.GetLessonCompletionsOfCourseAndUser)
	sMux.HandleFunc("POST /api/courses/{courseId}/completions", lessonCompletionController.CreateLessonCompletion)
	sMux.HandleFunc("OPTIONS /api/courses/{courseId}/completions", handleOptions)
	sMux.HandleFunc("GET /api/courses/completions", lessonCompletionController.GetAllLessonCompletionsOfUser)

	// users
	sMux.HandleFunc("OPTIONS /api/users", handleOptions)
	sMux.HandleFunc("POST /api/users", userController.CreateUser)

	sMux.HandleFunc("OPTIONS /api/users/login", handleOptions)
	sMux.HandleFunc("POST /api/users/login", userController.LoginAndReturnJwtToken)

	sMux.HandleFunc("OPTIONS /api/users/{userId}", handleOptions)
	sMux.HandleFunc("GET /api/users/{userId}", userController.GetUser)
	sMux.HandleFunc("PATCH /api/users/{userId}", userController.UpdateUser)

	// exams
	sMux.HandleFunc("GET /api/exams/{examId}", examController.GetExam)
	sMux.HandleFunc("POST /api/exams/{examId}/completions", examController.CompleteExam)
	sMux.HandleFunc("OPTIONS /api/exams/{examId}/completions", handleOptions)
	sMux.HandleFunc("GET /api/exams", examController.GetExamIds)
	sMux.HandleFunc("GET /api/exams/{examId}/completions", examController.GetCompletedExam)
	return sMux
}

func handleOptions(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.WriteHeader(http.StatusOK)
}

var _ http.Handler = (*PanicRecoveryMiddleware)(nil)

type PanicRecoveryMiddleware struct {
	Handler http.Handler
}

func (p *PanicRecoveryMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 2048)
			n := runtime.Stack(buf, false)
			buf = buf[:n]
			slog.Error("panic occurred", "error stack", string(buf))
			error_handling.WriteErrorDetailResponse(w, error_handling.ErrPanic)
		}
	}()
	p.Handler.ServeHTTP(w, r)
}

var _ http.Handler = (*CorsMiddleware)(nil)

type CorsMiddleware struct {
	Handler http.Handler
}

func (c *CorsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.Handler.ServeHTTP(w, r)
	cors := os.Getenv("PHBA_CORS")
	w.Header().Set("Access-Control-Allow-Origin", cors)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
