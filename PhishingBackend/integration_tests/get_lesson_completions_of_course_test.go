//go:build integration

package integration_tests

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestLessonCompletionsOfACourseCanBeRetrieved(t *testing.T) {
	// given
	user := createUser(t)
	jwtToken := getJwtTokenForUser(t, user)

	courseId1 := uuid.New()
	course1Lessons := []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}
	courseId2 := uuid.New()
	course2Lesson := uuid.New()

	for _, lesson := range course1Lessons {
		createSpecificLessonCompletion(t, jwtToken, courseId1, lesson)
	}
	createSpecificLessonCompletion(t, jwtToken, courseId2, course2Lesson)

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/courses/"+courseId2.String()+"/completions", nil)
	req.Header.Set("Authorization", "Bearer "+jwtToken)

	// when
	resp, err := http.DefaultClient.Do(req)

	// then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var lessons []uuid.UUID
	err = json.NewDecoder(resp.Body).Decode(&lessons)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(lessons))
	assert.Equal(t, course2Lesson, lessons[0])
}
