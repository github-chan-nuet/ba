package integration_tests

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestOptionsEndpoints(t *testing.T) {
	// given
	tests := []struct {
		url  string
		want []string
	}{{
		url:  "/api/health",
		want: []string{http.MethodGet, http.MethodHead, http.MethodOptions},
	}, {
		url:  "/api/courses/" + uuid.NewString() + "/completions",
		want: []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions},
	}, {
		url:  "/api/courses/completions",
		want: []string{http.MethodGet, http.MethodHead, http.MethodOptions},
	}, {
		url:  "/api/users",
		want: []string{http.MethodPost, http.MethodHead, http.MethodOptions},
	}, {
		url:  "/api/users/login",
		want: []string{http.MethodPost, http.MethodHead, http.MethodOptions},
	}, {
		url:  "/api/users/" + uuid.NewString(),
		want: []string{http.MethodGet, http.MethodPatch, http.MethodHead, http.MethodOptions},
	}, {
		url:  "/api/exams",
		want: []string{http.MethodGet, http.MethodHead, http.MethodOptions},
	}, {
		url:  "/api/exams/" + uuid.NewString(),
		want: []string{http.MethodGet, http.MethodHead, http.MethodOptions},
	}, {
		url:  "/api/exams/" + uuid.NewString() + "/completions",
		want: []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions},
	}, {
		url:  "/api/phishing-simulation/runs/" + uuid.NewString(),
		want: []string{http.MethodGet, http.MethodHead, http.MethodOptions},
	}}
	for _, tt := range tests {
		t.Run("Options "+tt.url, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodOptions, ts.URL+tt.url, nil)

			// when
			resp, err := http.DefaultClient.Do(req)

			// then
			assert.NoError(t, err)
			assert.ElementsMatch(t, tt.want, getMethods(resp.Header.Get("Allow")))
			assert.ElementsMatch(t, tt.want, getMethods(resp.Header.Get("Access-Control-Allow-Methods")))
			assert.Equal(t, "86400", resp.Header.Get("Access-Control-Max-Age"))
		})
	}
}

func getMethods(s string) []string {
	return strings.Split(s, ", ")
}
