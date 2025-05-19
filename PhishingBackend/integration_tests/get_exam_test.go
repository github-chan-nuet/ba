//go:build integration

package integration_tests

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/domain_model"
	"testing"
)

func TestExamCanBeGet(t *testing.T) {
	// given
	exam := &domain_model.Exam{
		ID: uuid.New(),
		Questions: []domain_model.ExamQuestion{
			{
				ID:       uuid.New(),
				Question: "A?",
				Answers: []domain_model.ExamQuestionAnswer{
					{
						ID:        uuid.New(),
						Answer:    "a",
						IsCorrect: false,
					},
					{
						ID:        uuid.New(),
						Answer:    "b",
						IsCorrect: true,
					},
				},
			},
			{
				ID:       uuid.New(),
				Question: "B?",
				Answers: []domain_model.ExamQuestionAnswer{
					{
						ID:        uuid.New(),
						Answer:    "a",
						IsCorrect: false,
					},
					{
						ID:        uuid.New(),
						Answer:    "b",
						IsCorrect: true,
					},
					{
						ID:        uuid.New(),
						Answer:    "c",
						IsCorrect: true,
					},
				},
			},
		},
	}
	createExam(t, exam)
	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/exams/"+exam.ID.String(), nil)

	// when
	resp, err := http.DefaultClient.Do(req)

	// then
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	var gotExam api.Exam
	err = json.NewDecoder(resp.Body).Decode(&gotExam)
	assert.NoError(t, err)
	assert.Equal(t, exam.ID, gotExam.Id)

	var qWith2Answer api.Question
	var qWith3Answer api.Question
	if len(gotExam.Questions[0].Answers) == 2 {
		qWith2Answer = gotExam.Questions[0]
		qWith3Answer = gotExam.Questions[1]
	} else {
		qWith2Answer = gotExam.Questions[1]
		qWith3Answer = gotExam.Questions[0]
	}

	wantQ2 := api.Question{
		Answers:  mapToApiAnswer(exam.Questions[0].Answers),
		Id:       exam.Questions[0].ID,
		Question: exam.Questions[0].Question,
		Type:     api.QuestionTypeSingleChoice,
	}
	assert.Equal(t, wantQ2, qWith2Answer)

	wantQ3 := api.Question{
		Answers:  mapToApiAnswer(exam.Questions[1].Answers),
		Id:       exam.Questions[1].ID,
		Question: exam.Questions[1].Question,
		Type:     api.QuestionTypeMultipleChoice,
	}
	assert.Equal(t, wantQ3, qWith3Answer)
}

func mapToApiAnswer(answers []domain_model.ExamQuestionAnswer) []api.Answer {
	dtoAnswers := make([]api.Answer, len(answers))
	for i, a := range answers {
		dtoAnswers[i] = api.Answer{
			Answer: a.Answer,
			Id:     a.ID,
		}
	}
	return dtoAnswers
}
