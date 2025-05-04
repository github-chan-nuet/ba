package error_handling

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"phishing_backend/internal/adapters/presentation/api"
)

func WriteErrorDetailResponse(w http.ResponseWriter, err error) {
	for errKey, prob := range problemMap {
		if errors.Is(err, errKey) {
			writeResponse(w, &prob)
			return
		}
	}
	writeResponse(w, &stdError)
}

func writeResponse(w http.ResponseWriter, prob *api.ProblemDetail) {
	slog.Error("writing error response", "type", prob.Type, "status", prob.Status, "title", prob.Title,
		"detail", prob.Detail, "instance", prob.Instance)
	w.WriteHeader(prob.Status)
	w.Header().Set("Content-Type", "application/problem+json")
	_ = json.NewEncoder(w).Encode(prob)
}
