package presentation

import (
	"log/slog"
	"net/http"
	"os"
	"phishing_backend/internal/application/services"
	"phishing_backend/internal/infrastructure/persistance"
	"phishing_backend/internal/infrastructure/presentation/controllers"
)

func SetupHttpServer() {
	sMux := http.NewServeMux()
	setupEndpoints(sMux)
	addr := os.Getenv("PHBA_WEBSERVER_ADDR")
	slog.Info("Web server listening...", "address", addr)
	err := http.ListenAndServe(addr, sMux)
	slog.Error("Web server stopped", "error", err)
	os.Exit(1)
}

func setupEndpoints(sMux *http.ServeMux) {
	userRepository := persistance.UserRepositoryImpl{}

	authenticator := services.AuthenticatorImpl{
		UserRepository: &userRepository,
	}
	userController := controllers.UserController{
		Authenticator: &authenticator,
		UserService: &services.UserServiceImpl{
			UserRepository: &userRepository,
		},
	}
	lessonCompletionController := controllers.LessonCompletionController{
		Authenticator: &authenticator,
		LessonCompletionService: &services.LessonCompletionServiceImpl{
			Repo: &persistance.LessonCompletionRepositoryImpl{},
		},
	}

	sMux.HandleFunc("GET /api/health", withCORS(controllers.GetHealth))

	sMux.HandleFunc("OPTIONS /api/courses/{courseId}/completions", withCORS(handleOptions))
	sMux.HandleFunc("POST /api/courses/{courseId}/completions", withCORS(lessonCompletionController.CreateLessonCompletion))

	sMux.HandleFunc("OPTIONS /api/users", withCORS(handleOptions))
	sMux.HandleFunc("POST /api/users", withCORS(userController.CreateUser))

	sMux.HandleFunc("OPTIONS /api/users/login", withCORS(handleOptions))
	sMux.HandleFunc("POST /api/users/login", withCORS(userController.LoginAndReturnJwtToken))

	sMux.HandleFunc("GET /api/users/{userId}", withCORS(userController.GetUser))
	sMux.HandleFunc("PATCH /api/users/{userId}", withCORS(userController.UpdateUser))
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Allow requests from any origin for now but tighten this is prod
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		next(w, r)
	}
}

func handleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.WriteHeader(http.StatusOK)
}
