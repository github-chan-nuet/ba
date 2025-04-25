//go:build integration

package integration_tests

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"phishing_backend/internal/infrastructure/presentation/api"
	"strings"
	"testing"
)

func createRandomEmail() string {
	return strings.ToLower(rand.Text()) + "@test.com"
}

func createUser(t *testing.T) *api.UserPostModel {
	reqBody := api.UserPostModel{
		Email:     createRandomEmail(),
		Password:  "password",
		Firstname: "John",
		Lastname:  "Doe",
	}
	marshal, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/users", bytes.NewReader(marshal))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("could not create user, expected %d, got %d", http.StatusCreated, resp.StatusCode)
		t.FailNow()
	}
	return &reqBody
}

func getJwtTokenForUser(t *testing.T, user *api.UserPostModel) string {
	reqBody := api.UserAuthentication{
		Email:    user.Email,
		Password: user.Password,
	}
	marshal, _ := json.Marshal(reqBody)
	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/users/login", bytes.NewReader(marshal))
	assert.NoError(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	binToken, _ := io.ReadAll(resp.Body)
	return string(binToken)
}

func getUserId(jwtToken string) uuid.UUID {
	claimString := strings.Split(jwtToken, ".")[1]
	decoded, _ := base64.StdEncoding.DecodeString(claimString)
	var result map[string]interface{}
	json.Unmarshal(decoded, &result)
	return uuid.MustParse(result["id"].(string))
}
