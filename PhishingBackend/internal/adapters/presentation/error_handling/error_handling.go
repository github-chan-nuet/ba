package error_handling

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/domain_model/validation"
)

func WriteErrorDetailResponse(w http.ResponseWriter, err error) {
	slog.Error("error handler received error", "err", err)
	var vErr *validation.ValidationError
	if errors.As(err, &vErr) {
		prob := toProblemDetail(vErr)
		writeResponse(w, prob)
		return
	}

	var pErr *PathVariableError
	if errors.As(err, &pErr) {
		prob := pErr.toProblemDetail()
		writeResponse(w, prob)
		return
	}

	for errKey, prob := range problemMap {
		if errors.Is(err, errKey) {
			writeResponse(w, &prob)
			return
		}
	}
	writeResponse(w, &stdProb)
}

func toProblemDetail(vErr *validation.ValidationError) *api.ProblemDetail {
	errs := make([]api.Error, 0, vErr.Len())
	for _, err := range vErr.Errors {
		errs = append(errs, api.Error{
			Detail:  string(err.Reason),
			Pointer: err.Field,
		})
	}

	return &api.ProblemDetail{
		Type:   createUrn("validation-error"),
		Title:  "Deine Anfrage ist nicht valide",
		Status: 422,
		Errors: &errs,
	}
}

func writeResponse(w http.ResponseWriter, prob *api.ProblemDetail) {
	slog.Error("writing error response", "type", prob.Type, "status", prob.Status, "title", prob.Title,
		"detail", prob.Detail, "instance", prob.Instance)
	w.WriteHeader(prob.Status)
	w.Header().Set("Content-Type", "application/problem+json")
	_ = json.NewEncoder(w).Encode(prob)
}
