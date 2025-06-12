package integration_tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"phishing_backend/internal/adapters/presentation/api"
	"testing"
)

func TestHttpCorsHeaders(t *testing.T) {
	// given
	reqBody := api.UserPostModel{
		Email:     createRandomEmail(),
		Password:  "password",
		Firstname: "John",
		Lastname:  "Doe",
	}
	marshal, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/users", bytes.NewReader(marshal))

	// when
	resp, err := http.DefaultClient.Do(req)

	// then
	assert.NoError(t, err)
	origin := resp.Header.Get("Access-Control-Allow-Origin")
	wantOrigin := os.Getenv("PHBA_CORS")
	assert.Equal(t, wantOrigin, origin)
	allowHeaders := resp.Header.Get("Access-Control-Allow-Headers")
	wantHeaders := "Content-Type, Authorization"
	assert.Equal(t, wantHeaders, allowHeaders)
}
