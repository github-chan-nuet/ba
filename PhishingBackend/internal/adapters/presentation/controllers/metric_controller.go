package controllers

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log/slog"
	"net/http"
)

func SetupPrometheus() {
	prometheus.MustRegister(collectors.NewGoCollector())
}

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	slog.Info("Get metrics called")
	promhttp.Handler().ServeHTTP(w, r)
}
