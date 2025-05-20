package presentation

import (
	"log/slog"
	"net/http"
	"os"
	"phishing_backend/internal/adapters/persistence"
	"phishing_backend/internal/adapters/presentation/controllers"
	"phishing_backend/internal/domain_services/services"
)

func SetupHttpServer() {
	sMux := NewServeMux()
	addr := os.Getenv("PHBA_WEBSERVER_ADDR")
	slog.Info("Web server listening...", "address", addr)
	err := http.ListenAndServe(addr, sMux)
	slog.Error("Web server stopped", "error", err)
	os.Exit(1)
}

func NewServeMux() *http.ServeMux {
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
	sMux.HandleFunc("GET /api/health", withCORS(controllers.GetHealth))

	// lesson completions
	sMux.HandleFunc("OPTIONS /api/courses/{courseId}/completions", withCORS(handleOptions))
	sMux.HandleFunc("POST /api/courses/{courseId}/completions", withCORS(lessonCompletionController.CreateLessonCompletion))

	sMux.HandleFunc("GET /api/courses/completions", withCORS(lessonCompletionController.GetAllLessonCompletionsOfUser))
	sMux.HandleFunc("GET /api/courses/{courseId}/completions", withCORS(lessonCompletionController.GetLessonCompletionsOfCourseAndUser))

	// users
	sMux.HandleFunc("OPTIONS /api/users", withCORS(handleOptions))
	sMux.HandleFunc("POST /api/users", withCORS(userController.CreateUser))

	sMux.HandleFunc("OPTIONS /api/users/login", withCORS(handleOptions))
	sMux.HandleFunc("POST /api/users/login", withCORS(userController.LoginAndReturnJwtToken))

	sMux.HandleFunc("GET /api/users/{userId}", withCORS(userController.GetUser))
	sMux.HandleFunc("PATCH /api/users/{userId}", withCORS(userController.UpdateUser))

	// exams
	sMux.HandleFunc("GET /api/exams/{examId}", withCORS(examController.GetExam))
	sMux.HandleFunc("POST /api/exams/{examId}/completions", withCORS(examController.CompleteExam))
	sMux.HandleFunc("OPTIONS /api/exams/{examId}/completions", withCORS(handleOptions))
	sMux.HandleFunc("GET /api/exams", withCORS(examController.GetExamIds))
	sMux.HandleFunc("GET /api/exams/{examId}/completions", withCORS(examController.GetCompletedExam))
	return sMux
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Allow requests from any origin for now but tighten this is prod
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		next(w, r)
	}
}

func handleOptions(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.WriteHeader(http.StatusOK)
}
