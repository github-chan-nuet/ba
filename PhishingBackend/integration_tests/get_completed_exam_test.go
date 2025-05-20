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
	wantCompQs := make([]api.CompletedQuestion, len(exam.Questions))
	i, j := 0, 0
	if compExam.Questions[0].Id != exam.Questions[0].ID {
		i = 1
	}
	wantCompQs[0] = api.CompletedQuestion{
		Answers:     *toAnswerWithCorrections(&exam.Questions[i].Answers),
		Id:          exam.Questions[i].ID,
		Question:    exam.Questions[i].Question,
		Type:        getQuestionType(&exam.Questions[i]),
		UserAnswers: exComp[j].Answers,
	}
	i = (i + 1) % 2
	j++
	wantCompQs[1] = api.CompletedQuestion{
		Answers:     *toAnswerWithCorrections(&exam.Questions[i].Answers),
		Id:          exam.Questions[i].ID,
		Question:    exam.Questions[i].Question,
		Type:        getQuestionType(&exam.Questions[i]),
		UserAnswers: exComp[j].Answers,
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
