package integration_tests

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"phishing_backend/internal/domain_model"
	"testing"
)

func TestExamIdsCanBeGet(t *testing.T) {
	// given
	exam := &domain_model.Exam{ID: uuid.New()}
	createExam(t, exam)
	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/exams", nil)

	// when
	resp, err := http.DefaultClient.Do(req)

	// then
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	var gotExamIds []uuid.UUID
	err = json.NewDecoder(resp.Body).Decode(&gotExamIds)
	assert.NoError(t, err)
	// at least one examId is returned. It could be that other test cases create exams as well
	assert.GreaterOrEqual(t, len(gotExamIds), 1)
	assert.Contains(t, gotExamIds, exam.ID)
}
