//go:build integration

package integration_tests

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/domain_model"
	"testing"
)

func TestLessonCanBeCompleted(t *testing.T) {
	// given
	user := createUser(t)
	jwtToken := getJwtTokenForUser(t, user)

	reqBody := api.Lesson{LessonId: uuid.New()}
	marshal, _ := json.Marshal(reqBody)
	url := ts.URL + "/api/courses/" + uuid.NewString() + "/completions"
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshal))
	req.Header.Set("Authorization", "Bearer "+jwtToken)

	// when
	resp, err := http.DefaultClient.Do(req)

	// then
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	var expGain api.ExperienceGain
	err = json.NewDecoder(resp.Body).Decode(&expGain)
	assert.NoError(t, err)
	assert.Equal(t, int64(domain_model.LessonCompletionGain), expGain.NewExperienceGained)
	assert.Equal(t, int64(domain_model.LessonCompletionGain), expGain.TotalExperience)
	assert.NotNil(t, expGain.NewLevel)
	assert.Equal(t, int64(2), *expGain.NewLevel)
}

func TestTheSameLessonCantBeCompletedTwice(t *testing.T) {
	t.SkipNow() // todo enable this test once the same lesson can't be inserted twice

	// given
	user := createUser(t)
	jwtToken := getJwtTokenForUser(t, user)

	reqBody := api.Lesson{LessonId: uuid.New()}
	marshal, _ := json.Marshal(reqBody)
	url := ts.URL + "/api/courses/" + uuid.NewString() + "/completions"
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshal))
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	// and lesson was already completed
	http.DefaultClient.Do(req)
	// and create new request for the same lesson completion
	req2, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshal))
	req2.Header.Set("Authorization", "Bearer "+jwtToken)

	// when
	resp, err := http.DefaultClient.Do(req2)

	// then
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
