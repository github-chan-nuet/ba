package web

import (
	"log/slog"
	"net/http"
	"os"
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
	sMux.HandleFunc("GET /api/health", getHealth)
	sMux.HandleFunc("POST /api/users/login", loginAndReturnJwtToken)
	sMux.HandleFunc("POST /api/courses/{courseId}/completions", createLessonCompletion)
}
