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
	"strings"
	"sync"
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
		UserRepo:          &userRepository,
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
	routes := NewMethodMap()

	// health
	routes.Add(sMux, "/api/health", http.MethodGet, controllers.GetHealth)

	// lesson completions
	routes.Add(sMux, "/api/courses/{courseId}/completions", http.MethodGet, lessonCompletionController.GetLessonCompletionsOfCourseAndUser)
	routes.Add(sMux, "/api/courses/{courseId}/completions", http.MethodPost, lessonCompletionController.CreateLessonCompletion)

	routes.Add(sMux, "/api/courses/completions", http.MethodGet, lessonCompletionController.GetAllLessonCompletionsOfUser)

	// users
	routes.Add(sMux, "/api/users", http.MethodPost, userController.CreateUser)

	routes.Add(sMux, "/api/users/login", http.MethodPost, userController.LoginAndReturnJwtToken)

	routes.Add(sMux, "/api/users/{userId}", http.MethodGet, userController.GetUser)
	routes.Add(sMux, "/api/users/{userId}", http.MethodPatch, userController.UpdateUser)

	// exams
	routes.Add(sMux, "/api/exams", http.MethodGet, examController.GetExams)

	routes.Add(sMux, "/api/exams/{examId}", http.MethodGet, examController.GetExam)

	routes.Add(sMux, "/api/exams/{examId}/completions", http.MethodGet, examController.GetCompletedExam)
	routes.Add(sMux, "/api/exams/{examId}/completions", http.MethodPost, examController.CompleteExam)
	return sMux
}

type MethodMap struct {
	mu        sync.RWMutex
	endpoints map[string]map[string]http.HandlerFunc
}

func NewMethodMap() *MethodMap {
	return &MethodMap{
		endpoints: make(map[string]map[string]http.HandlerFunc),
	}
}

func (m *MethodMap) Add(mux *http.ServeMux, path, method string, handler http.HandlerFunc) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.endpoints[path]; !exists {
		m.endpoints[path] = make(map[string]http.HandlerFunc)
	}

	m.endpoints[path][method] = handler

	// Wrap the handler only once per path
	if len(m.endpoints[path]) == 1 {
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			m.mu.RLock()
			methods := m.endpoints[path]
			m.mu.RUnlock()

			// Build Allow header
			allowedMethods := make([]string, 0, len(methods))
			for method := range methods {
				allowedMethods = append(allowedMethods, method)
			}
			allowedMethods = append(allowedMethods, http.MethodOptions)

			if r.Method == http.MethodOptions {
				w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(allowedMethods, ", "))
				w.Header().Set("Access-Control-Max-Age", "86400")
				w.WriteHeader(http.StatusNoContent)
				return
			}

			if h, ok := methods[r.Method]; ok {
				h(w, r)
			} else {
				w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		})
	}
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
	cors := os.Getenv("PHBA_CORS")
	w.Header().Set("Access-Control-Allow-Origin", cors)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Handler.ServeHTTP(w, r)
}
