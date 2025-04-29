//go:build integration

package integration_tests

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"phishing_backend/internal/adapters/presentation/api"
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
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("could not create user, expected %d, got %d", http.StatusCreated, resp.StatusCode)
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
	token, _, _ := new(jwt.Parser).ParseUnverified(jwtToken, jwt.MapClaims{})
	return uuid.MustParse(token.Claims.(jwt.MapClaims)["id"].(string))
}

func createLessonCompletion(t *testing.T, jwtToken string) *api.ExperienceGain {
	reqBody := api.Lesson{LessonId: uuid.New()}
	marshal, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/courses/"+uuid.NewString()+"/completions", bytes.NewReader(marshal))
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusCreated)
	var expGain api.ExperienceGain
	if err := json.NewDecoder(resp.Body).Decode(&expGain); err != nil {
		t.Fatal("failed to decode response:", err)
	}
	return &expGain
}
