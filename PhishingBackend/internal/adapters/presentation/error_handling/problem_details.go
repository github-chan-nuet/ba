package error_handling

import (
	"errors"
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	. "phishing_backend/internal/domain_services/services"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidBody  = errors.New("invalid HTTP request body")
	ErrPanic        = errors.New("panic occurred")
	invalidJwtToken = api.ProblemDetail{
		Type:   createUrn("invalid-jwt-token"),
		Title:  "Dein JWT-Token ist ungültig",
		Status: 401,
	}
	stdProb = api.ProblemDetail{
		Type:   createUrn("generic-error"),
		Title:  "Uuuuupps, etwas lief schief",
		Status: 500,
	}
	problemMap = map[error]api.ProblemDetail{
		repositories.ErrEmailAlreadyUsed: {
			Type:   createUrn("email-already-used"),
			Title:  "Diese Email wird bereits verwendet",
			Status: 422,
		},
		ErrUnauthorized: {
			Type:   createUrn("unauthorized"),
			Title:  "Du bist nicht berechtigt, diese Operation auszuführen",
			Status: 403,
		},
		ErrNoAuthToken: {
			Type:   createUrn("not-authenticated"),
			Title:  "You are not authenticated",
			Status: 401,
			Detail: ptr("No Authorization header is present, please login"),
		},
		ErrInvalidAuthHeader:          invalidJwtToken,
		ErrInvalidTokenSignMethod:     invalidJwtToken,
		ErrInvalidToken:               invalidJwtToken,
		ErrInvalidTokenClaimStructure: invalidJwtToken,
		ErrInvalidBody: {
			Type:   createUrn("invalid-http-request-body"),
			Title:  "Deine Anfrage ist invalide",
			Status: 422,
		},
		ErrQuestionNotExisting: {
			Type:   createUrn("question-does-not-exist"),
			Title:  "Die Testfrage existiert nicht",
			Status: 422,
		},
		ErrPanic: stdProb,
	}
)

func createUrn(identifier string) string {
	return "urn:securaware:error:" + identifier
}

func ptr(s string) *string {
	return &s
}
