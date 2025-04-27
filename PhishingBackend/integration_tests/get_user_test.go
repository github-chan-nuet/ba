//go:build integration

package integration_tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"phishing_backend/internal/infrastructure/presentation/api"
	"testing"
)

func TestUserInformationCanBeRetrieved(t *testing.T) {
	// given
	user := createUser(t)
	jwtToken := getJwtTokenForUser(t, user)
	userId := getUserId(jwtToken)

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/users/"+userId.String(), nil)
	req.Header.Set("Authorization", "Bearer "+jwtToken)

	// when
	resp, err := http.DefaultClient.Do(req)

	// then
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	var gotUser api.User
	err = json.NewDecoder(resp.Body).Decode(&gotUser)
	assert.NoError(t, err)
	assert.Equal(t, user.Firstname, *gotUser.Firstname)
	assert.Equal(t, user.Lastname, *gotUser.Lastname)
	assert.Equal(t, user.Email, *gotUser.Email)
	assert.Equal(t, 1, *gotUser.Level)
	assert.Equal(t, 0, *gotUser.TotalExperience)
}

func TestLevelAndExperienceCanBeRetrieved(t *testing.T) {
	// given
	user := createUser(t)
	jwtToken := getJwtTokenForUser(t, user)
	userId := getUserId(jwtToken)
	expGain := createLessonCompletion(t, jwtToken)

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/users/"+userId.String(), nil)
	req.Header.Set("Authorization", "Bearer "+jwtToken)

	// when
	resp, err := http.DefaultClient.Do(req)

	// then
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	var gotUser api.User
	err = json.NewDecoder(resp.Body).Decode(&gotUser)
	assert.NoError(t, err)
	assert.Equal(t, *expGain.NewLevel, int64(*gotUser.Level))
	assert.Equal(t, expGain.TotalExperience, int64(*gotUser.TotalExperience))
}
