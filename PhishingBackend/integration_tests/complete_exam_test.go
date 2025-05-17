//go:build integration

package integration_tests

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"math"
	"net/http"
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/domain_model"
	"testing"
)

func Test(t *testing.T) {
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

	// when
	resp, err := http.DefaultClient.Do(req)

	// then
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	var gotExpGain api.ExperienceGain
	err = json.NewDecoder(resp.Body).Decode(&gotExpGain)
	assert.NoError(t, err)
	wantExp := int(0.17 * domain_model.ExamCompletionGain)
	wantLvl := calcLevel(wantExp)
	assert.Equal(t, wantLvl, *gotExpGain.NewLevel)
	assert.Equal(t, wantExp, gotExpGain.TotalExperience)
	assert.Equal(t, wantExp, gotExpGain.NewExperienceGained)
}

func calcLevel(totalExperience int) int {
	levelAsFloat := float64(1) + (math.Log((float64(totalExperience)/200)+1) / math.Log(1.5))
	lvl := int(math.Floor(levelAsFloat))
	return lvl
}
