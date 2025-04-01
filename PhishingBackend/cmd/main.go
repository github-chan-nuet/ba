package main

import (
	"log/slog"
	"net/http"
	"os"
	"phishing_backend/internal/email"
)

func main() {
	setupDefaultLogger()
	// listEnvironmentVariables()

	err := email.SendEmail()
	if err != nil {
		slog.Error(err.Error())
	}
}

func setupHttpServer() {
	sMux := http.NewServeMux()
	sMux.HandleFunc("GET /api/health", getHealth)
	addr, ok := os.LookupEnv("PHBA_WEBSERVER_ADDR")
	if !ok {
		addr = "localhost:8080"
	}
	slog.Info("Web server listening...", "address", addr)
	err := http.ListenAndServe(addr, sMux)
	slog.Error("Web server stopped", "error", err)
	os.Exit(1)
}

func getHealth(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func setupDefaultLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func listEnvironmentVariables() {
	envs := os.Environ()
	for _, env := range envs {
		slog.Info(env)
	}
}
