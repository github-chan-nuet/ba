package main

import (
	"log/slog"
	"net/http"
	"os"
	_ "phishing_backend/internal/domain/db" // include package so that init function is called
)

func main() {
	// listEnvironmentVariables()
	//err := email.SendEmail()
	//if err != nil {
	//	slog.Error(err.Error())
	//}
	setupHttpServer()
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}

func init() {
	setupDefaultLogger()
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
