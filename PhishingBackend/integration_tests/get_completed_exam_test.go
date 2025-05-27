//go:build integration

package integration_tests

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/domain_model"
	"testing"
	"time"
)

func TestCanRetrieveCompletedExam(t *testing.T) {
	// given
	exam := &domain_model.Exam{
		ID: uuid.New(),
		Questions: []domain_model.ExamQuestion{
			{
				ID:       uuid.New(),
				Question: "1 + 1 = ?",
				Answers: []domain_model.ExamQuestionAnswer{
					{
						ID:        uuid.New(),
						Answer:    "1",
						IsCorrect: false,
					},
					{
						ID:        uuid.New(),
						Answer:    "2",
						IsCorrect: true,
					},
				},
			},
			{
				ID:       uuid.New(),
				Question: "1 = ?",
				Answers: []domain_model.ExamQuestionAnswer{
					{
						ID:        uuid.New(),
						Answer:    "2",
						IsCorrect: false,
					},
					{
						ID:        uuid.New(),
						Answer:    "1",
						IsCorrect: true,
					},
					{
						ID:        uuid.New(),
						Answer:    "1 + 0",
						IsCorrect: true,
					},
				},
			},
		},
	}
	createExam(t, exam)
	exComp := api.ExamCompletion{
		{
			QuestionId: exam.Questions[0].ID,
			Answers:    []uuid.UUID{exam.Questions[0].Answers[0].ID},
		},
		{
			QuestionId: exam.Questions[1].ID,
			Answers: []uuid.UUID{
				exam.Questions[1].Answers[0].ID,
				exam.Questions[1].Answers[1].ID,
			},
		},
	}

	user := createUser(t)
	jwtToken := getJwtTokenForUser(t, user)
	marshal, _ := json.Marshal(exComp)
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/exams/"+exam.ID.String()+"/completions", bytes.NewReader(marshal))
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	_, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	req2, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/exams/"+exam.ID.String()+"/completions", nil)
	req2.Header.Set("Authorization", "Bearer "+jwtToken)

	// when
	resp, err := http.DefaultClient.Do(req2)

	// then
	var compExam api.CompletedExam
	json.NewDecoder(resp.Body).Decode(&compExam)
	wantCompletedAt := time.Now().UTC().Truncate(24 * time.Hour)
	assert.Equal(t, wantCompletedAt, compExam.CompletedAt.Time)
	wantCompQs := [...]api.CompletedQuestion{
		{
			Answers:     *toAnswerWithCorrections(&exam.Questions[0].Answers),
			Id:          exam.Questions[0].ID,
			Question:    exam.Questions[0].Question,
			Type:        getQuestionType(&exam.Questions[0]),
			UserAnswers: exComp[0].Answers,
		},
		{
			Answers:     *toAnswerWithCorrections(&exam.Questions[1].Answers),
			Id:          exam.Questions[1].ID,
			Question:    exam.Questions[1].Question,
			Type:        getQuestionType(&exam.Questions[1]),
			UserAnswers: exComp[1].Answers,
		},
	}
	assert.ElementsMatch(t, wantCompQs, compExam.Questions)
}

func getQuestionType(q *domain_model.ExamQuestion) api.CompletedQuestionType {
	numCorrect := 0
	for _, a := range q.Answers {
		if a.IsCorrect {
			numCorrect++
		}
	}
	if numCorrect == 1 {
		return api.CompletedQuestionTypeSingleChoice
	}
	return api.CompletedQuestionTypeMultipleChoice
}

func toAnswerWithCorrections(answers *[]domain_model.ExamQuestionAnswer) *[]api.AnswerWithCorrection {
	dtos := make([]api.AnswerWithCorrection, len(*answers))
	for i, answer := range *answers {
		dtos[i] = *toAnswerWithCorrection(&answer)
	}
	return &dtos
}

func toAnswerWithCorrection(answer *domain_model.ExamQuestionAnswer) *api.AnswerWithCorrection {
	dto := api.AnswerWithCorrection{
		Answer:    answer.Answer,
		Id:        answer.ID,
		IsCorrect: answer.IsCorrect,
	}
	return &dto
}
