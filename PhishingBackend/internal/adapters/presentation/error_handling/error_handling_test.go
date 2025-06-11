package error_handling

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/domain_model/validation"
	"testing"
)

func TestErrorsAreReturnedAsErrorDetails(t *testing.T) {
	for err, problemDetail := range problemMap {
		t.Run(problemDetail.Type, func(t *testing.T) {
			// given
			recorder := httptest.NewRecorder()

			// when
			WriteErrorDetailResponse(recorder, err)

			// then
			resp := recorder.Result()
			assert.Equal(t, problemDetail.Status, resp.StatusCode)
			body, _ := io.ReadAll(resp.Body)
			wantBody, _ := json.Marshal(problemDetail)
			assert.Equal(t, string(wantBody)+"\n", string(body))
			assert.Equal(t, "application/problem+json", recorder.Header().Get("Content-Type"))
		})
	}
}

func TestUnknownError(t *testing.T) {
	// given
	recorder := httptest.NewRecorder()
	err := errors.New("unknown error")

	// when
	WriteErrorDetailResponse(recorder, err)

	// then
	resp := recorder.Result()
	assert.Equal(t, stdProb.Status, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	wantBody, _ := json.Marshal(stdProb)
	assert.Equal(t, string(wantBody)+"\n", string(body))
	assert.Equal(t, "application/problem+json", recorder.Header().Get("Content-Type"))
}

func TestValidationError(t *testing.T) {
	// given
	recorder := httptest.NewRecorder()
	vErr := validation.NewValidationError()
	vErr.Add("#/firstname", validation.Mandatory)

	// when
	WriteErrorDetailResponse(recorder, vErr)

	// then
	wantProblemDetail := api.ProblemDetail{
		Type:   createUrn("validation-error"),
		Title:  "Deine Anfrage ist nicht valide",
		Status: 422,
		Errors: &[]api.Error{{
			Detail:  string(validation.Mandatory),
			Pointer: "#/firstname",
		}},
	}
	resp := recorder.Result()
	assert.Equal(t, wantProblemDetail.Status, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	wantBody, _ := json.Marshal(wantProblemDetail)
	assert.Equal(t, string(wantBody)+"\n", string(body))
	assert.Equal(t, "application/problem+json", recorder.Header().Get("Content-Type"))
}

func TestPathVariableError(t *testing.T) {
	// given
	recorder := httptest.NewRecorder()
	err := PathVariableError{}

	// when
	WriteErrorDetailResponse(recorder, &err)

	// then
	wantProblemDetail := api.ProblemDetail{
		Type:   createUrn("path-variable-error"),
		Title:  fmt.Sprintf("Deine Anfrage ist nicht valide. Die Variable %s in der URL fehlt oder ist falsch", err.Name),
		Status: 422,
	}
	resp := recorder.Result()
	assert.Equal(t, wantProblemDetail.Status, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	wantBody, _ := json.Marshal(wantProblemDetail)
	assert.Equal(t, string(wantBody)+"\n", string(body))
	assert.Equal(t, "application/problem+json", recorder.Header().Get("Content-Type"))
}
