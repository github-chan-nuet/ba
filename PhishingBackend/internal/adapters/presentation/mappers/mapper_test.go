package mappers

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/domain_model"
	"testing"
)

// ----- ToUserPatchDto -----
func TestToUserPatchDto(t *testing.T) {
	// given
	patch := api.UserPatchModel{
		Email:     ptr("a"),
		Firstname: ptr("b"),
		Lastname:  ptr("c"),
		Password:  ptr("d"),
	}

	// when
	dto := ToUserPatchDto(patch)

	// then
	assert.Equal(t, patch.Email, dto.Email)
	assert.Equal(t, patch.Firstname, dto.Firstname)
	assert.Equal(t, patch.Lastname, dto.Lastname)
	assert.Equal(t, patch.Password, dto.Password)
}

// ----- ToUserPostDto -----
func TestToUserPostDto(t *testing.T) {
	// given
	post := api.UserPostModel{
		Email:     "a",
		Firstname: "b",
		Lastname:  "c",
		Password:  "d",
	}

	// when
	dto := ToUserPostDto(post)

	// then
	assert.Equal(t, post.Email, dto.Email)
	assert.Equal(t, post.Firstname, dto.Firstname)
	assert.Equal(t, post.Lastname, dto.Lastname)
	assert.Equal(t, post.Password, dto.Password)
}

// ----- ToQuestionCompletionDto -----
func TestToQuestionCompletionDto(t *testing.T) {
	// given
	qs := []api.QuestionCompletion{
		{
			Answers:    []uuid.UUID{uuid.New()},
			QuestionId: uuid.New(),
		},
	}

	// when
	dto := ToQuestionCompletionDto(&qs)

	// then
	assert.Len(t, *dto, 1)
	assert.Equal(t, qs[0].QuestionId, (*dto)[0].QuestionId)
	assert.Equal(t, qs[0].Answers[0], (*dto)[0].Answers[0])
}

// ----- ToApiExpGain -----
func TestToApiExpGain(t *testing.T) {
	// given
	expGain := domain_model.ExperienceGain{
		NewExperienceGained: 1,
		TotalExperience:     2,
		NewLevel:            ptr(3),
	}

	// when
	apiExpGain := ToApiExpGain(&expGain)

	// then
	assert.Equal(t, expGain.NewExperienceGained, apiExpGain.NewExperienceGained)
	assert.Equal(t, expGain.TotalExperience, apiExpGain.TotalExperience)
	assert.Equal(t, *expGain.NewLevel, *apiExpGain.NewLevel)
}

func ptr[v any](s v) *v {
	return &s
}
