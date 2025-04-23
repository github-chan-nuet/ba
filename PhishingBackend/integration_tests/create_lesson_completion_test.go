//go:build integration

package integration_tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"phishing_backend/internal/infrastructure/presentation"
	"phishing_backend/internal/infrastructure/presentation/api"
	"testing"
)

// https://medium.com/insiderengineering/integration-test-in-golang-899412b7e1bf
func Test(t *testing.T) {
	ts := httptest.NewServer(presentation.NewServeMux())
	defer ts.Close()

	// given
	reqBody := api.UserPostModel{}
	marshal, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/users", bytes.NewReader(marshal))

	// when
	resp, err := http.DefaultClient.Do(req)

	// then
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusCreated)
}
