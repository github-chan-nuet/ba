//go:build integration

package integration_tests

import (
	"encoding/json"
	"net/http"
	"phishing_backend/internal/domain_model"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestExamsCanBeGet(t *testing.T) {
	// given
	exam := &domain_model.Exam{ID: uuid.New(), Questions: []domain_model.ExamQuestion{}}
	createExam(t, exam)
	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/exams", nil)

	// when
	resp, err := http.DefaultClient.Do(req)

	// then
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	var gotExams []domain_model.Exam
	err = json.NewDecoder(resp.Body).Decode(&gotExams)
	assert.NoError(t, err)
	// at least one examId is returned. It could be that other test cases create exams as well
	assert.GreaterOrEqual(t, len(gotExams), 1)
	assert.Contains(t, gotExams, *exam)
}
