package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"phishing_backend/internal/adapters/presentation/error_handling"
)

func writeJsonResponse(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func getPathVariable(r *http.Request, pName string) (uuid.UUID, error) {
	pVar, err := uuid.Parse(r.PathValue(pName))
	if err != nil {
		return uuid.Nil, &error_handling.PathVariableError{Name: pName}
	}
	return pVar, nil
}
