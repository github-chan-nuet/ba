package integration_tests

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"phishing_backend/internal/adapters/presentation/api"
	"testing"
)

func TestAllLessonCompletionsCanBeRetrieved(t *testing.T) {
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

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/courses/completions", nil)
	req.Header.Set("Authorization", "Bearer "+jwtToken)

	// when
	resp, err := http.DefaultClient.Do(req)

	// then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var completions []api.CourseCompletion
	err = json.NewDecoder(resp.Body).Decode(&completions)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(completions))
	for _, completion := range completions {
		if completion.CourseId == courseId1 {
			assert.ElementsMatch(t, course1Lessons, completion.CompletedLessons)
		} else {
			assert.ElementsMatch(t, []uuid.UUID{course2Lesson}, completion.CompletedLessons)
		}
	}
}
