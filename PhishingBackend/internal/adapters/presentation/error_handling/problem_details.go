package error_handling

import (
	"errors"
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"phishing_backend/internal/domain_services/services"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	invalidJwtToken = api.ProblemDetail{
		Type:   createUrn("invalid-jwt-token"),
		Title:  "Your JWT token is invalid",
		Status: 400,
	}
	problemMap = map[error]api.ProblemDetail{
		repositories.ErrEmailAlreadyUsed: {
			Type:   createUrn("email-already-used"),
			Title:  "Email is already used and may not be used twice",
			Status: 422,
		},
		ErrUnauthorized: {
			Type:   createUrn("unauthorized"),
			Title:  "You are not authorized to perform this operation",
			Status: 403,
		},
		services.ErrNoAuthToken: {
			Type:   createUrn("not-authenticated"),
			Title:  "You are not authenticated",
			Status: 401,
			Detail: ptr("No Authorization header is present, please login"),
		},
		services.ErrInvalidAuthHeader:          invalidJwtToken,
		services.ErrInvalidTokenSignMethod:     invalidJwtToken,
		services.ErrInvalidToken:               invalidJwtToken,
		services.ErrInvalidTokenClaimStructure: invalidJwtToken,
	}
	stdError = api.ProblemDetail{
		Type:   createUrn("generic-error"),
		Title:  "An internal server error has occurred",
		Status: 500,
	}
)

func createUrn(identifier string) string {
	return "urn:securaware:error:" + identifier
}

func ptr(s string) *string {
	return &s
}
