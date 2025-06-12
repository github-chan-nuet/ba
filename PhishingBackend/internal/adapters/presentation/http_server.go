package presentation

import (
	"log/slog"
	"net/http"
	"os"
	"phishing_backend/internal/adapters"
	"phishing_backend/internal/adapters/presentation/controllers"
	"phishing_backend/internal/adapters/presentation/error_handling"
	"runtime"
	"strings"
)

func SetupHttpServer(d *adapters.Dependencies) {
	handler := NewHttpHandler(d)
	addr := os.Getenv("PHBA_WEBSERVER_ADDR")
	slog.Info("Web server listening...", "address", addr)
	err := http.ListenAndServe(addr, handler)
	slog.Error("Web server stopped", "error", err)
	os.Exit(1)
}

func NewHttpHandler(d *adapters.Dependencies) http.Handler {
	sMux := newSecurawareEndpointHandler(d)
	cors := CorsMiddleware{Handler: sMux}
	return &PanicRecoveryMiddleware{Handler: &cors}
}

func newSecurawareEndpointHandler(d *adapters.Dependencies) http.Handler {
	// controllers
	userController := controllers.UserController{
		Authenticator:     d.Authenticator,
		UserService:       d.UserService,
		UserRepo:          d.UserRepository,
		ExperienceService: d.ExperienceService,
	}
	lessonCompletionController := controllers.LessonCompletionController{
		Authenticator:              d.Authenticator,
		LessonCompletionService:    d.LessonCompletionService,
		ExperienceService:          d.ExperienceService,
		LessonCompletionRepository: d.LessonCompletionRepository,
	}
	examController := controllers.ExamController{
		Authenticator:         d.Authenticator,
		ExamRepository:        d.ExamRepository,
		ExamCompRepo:          d.ExamCompletionRepository,
		ExamCompletionService: d.ExamCompletionService,
	}
	phishingSimController := controllers.PhishingSimulationController{
		PhishingRunService:           d.PhishingRunService,
		PhishingSimulationRepository: d.PhishingSimulationRepository,
	}

	routes := NewSmuxCreator()

	// health
	routes.Add("/api/health", http.MethodGet, controllers.GetHealth)

	// lesson completions
	routes.Add("/api/courses/{courseId}/completions", http.MethodGet, lessonCompletionController.GetLessonCompletionsOfCourseAndUser)
	routes.Add("/api/courses/{courseId}/completions", http.MethodPost, lessonCompletionController.CreateLessonCompletion)

	routes.Add("/api/courses/completions", http.MethodGet, lessonCompletionController.GetAllLessonCompletionsOfUser)

	// users
	routes.Add("/api/users", http.MethodPost, userController.CreateUser)

	routes.Add("/api/users/login", http.MethodPost, userController.LoginAndReturnJwtToken)

	routes.Add("/api/users/{userId}", http.MethodGet, userController.GetUser)
	routes.Add("/api/users/{userId}", http.MethodPatch, userController.UpdateUser)

	// exams
	routes.Add("/api/exams", http.MethodGet, examController.GetExams)

	routes.Add("/api/exams/{examId}", http.MethodGet, examController.GetExam)

	routes.Add("/api/exams/{examId}/completions", http.MethodGet, examController.GetCompletedExam)
	routes.Add("/api/exams/{examId}/completions", http.MethodPost, examController.CompleteExam)

	// phishing simulation
	routes.Add("/api/phishing-simulation/runs/{phishingSimulationRunId}", http.MethodGet, phishingSimController.GetRun)

	return routes.BuildWithOptionEndpoints()
}

type SmuxCreator struct {
	// map of paths of methods -> handlers
	endpoints map[string]map[string]http.HandlerFunc
}

func NewSmuxCreator() *SmuxCreator {
	return &SmuxCreator{
		endpoints: make(map[string]map[string]http.HandlerFunc),
	}
}

func (m *SmuxCreator) Add(path, method string, handler http.HandlerFunc) {
	if _, exists := m.endpoints[path]; !exists {
		m.endpoints[path] = make(map[string]http.HandlerFunc)
	}
	m.endpoints[path][method] = handler
}

func (m *SmuxCreator) addOptionsEndpoints() {
	for path, methodMap := range m.endpoints {
		allowedMethods := make([]string, 0, len(methodMap)+2)
		for method := range methodMap {
			allowedMethods = append(allowedMethods, method)
		}
		allowedMethods = append(allowedMethods, http.MethodOptions, http.MethodHead)
		m.endpoints[path][http.MethodOptions] = m.createOptionsEndpoint(allowedMethods)
	}
}

func (m *SmuxCreator) createOptionsEndpoint(allowedMethods []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		methods := strings.Join(allowedMethods, ", ")
		w.Header().Set("Allow", methods)
		w.Header().Set("Access-Control-Allow-Methods", methods)
		w.Header().Set("Access-Control-Max-Age", "86400") // one day in seconds
		w.WriteHeader(http.StatusNoContent)
	}
}

func (m *SmuxCreator) BuildWithOptionEndpoints() *http.ServeMux {
	m.addOptionsEndpoints()
	sMux := http.NewServeMux()
	for path, methodMap := range m.endpoints {
		for method, handler := range methodMap {
			sMux.HandleFunc(method+" "+path, handler)
		}
	}
	return sMux
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
