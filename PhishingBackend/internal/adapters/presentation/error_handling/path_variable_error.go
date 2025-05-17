package error_handling

import (
	"fmt"
	"phishing_backend/internal/adapters/presentation/api"
)

type PathVariableError struct {
	Name string
}

func (p *PathVariableError) Error() string {
	return fmt.Sprintf("Path variable '%s' missing or incorrect", p.Name)
}

func (p *PathVariableError) toProblemDetail() *api.ProblemDetail {
	return &api.ProblemDetail{
		Type:   createUrn("path-variable-error"),
		Title:  fmt.Sprintf("Deine Anfrage ist nicht valide. Die Variable %s in der URL fehlt oder ist falsch", p.Name),
		Status: 422,
	}
}
