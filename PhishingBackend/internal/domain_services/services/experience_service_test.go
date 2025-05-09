package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"math"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"testing"
)

type MockExamCompRepo struct {
	repositories.ExamCompletionRepository
}

func (m *MockExamCompRepo) GetScores(userId uuid.UUID) ([]int, error) {
	return []int{}, nil
}

// ----- GetEntireExperience -----
func TestExperienceCalculationIsCorrect(t *testing.T) {
	// given
	m := repositories.NewMockLessonCompletionRepository(gomock.NewController(t))
	sut := ExperienceServiceImpl{LessonCompRepo: m, ExamCompRepo: &MockExamCompRepo{}}

	tests := []struct {
		name          string
		nLesson       int
		wantTotalExp  int
		wantCalcLevel int
	}{
		{
			name:          "No lessons completed",
			nLesson:       0,
			wantTotalExp:  0,
			wantCalcLevel: 1,
		},
		{
			name:          "One lesson completed",
			nLesson:       1,
			wantTotalExp:  domain_model.LessonCompletionGain,
			wantCalcLevel: 2,
		},
		{
			name:          "10 lessons completed",
			nLesson:       10,
			wantTotalExp:  10 * domain_model.LessonCompletionGain,
			wantCalcLevel: calcLevel(10 * domain_model.LessonCompletionGain),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.EXPECT().CountForUser(gomock.Any()).Return(tt.nLesson, nil)

			// when
			exp, err := sut.GetEntireExperience(uuid.New())

			// then
			assert.NoError(t, err)
			assert.Equal(t, tt.wantTotalExp, exp.TotalExperience)
			assert.Equal(t, tt.wantCalcLevel, exp.Level)
		})
	}
}

func TestExperienceCalculationReturnsErrorFromRepository(t *testing.T) {
	// given
	m := repositories.NewMockLessonCompletionRepository(gomock.NewController(t))
	givenErr := errors.New("some error")
	m.EXPECT().CountForUser(gomock.Any()).Return(0, givenErr)
	sut := ExperienceServiceImpl{LessonCompRepo: m}

	// when
	exp, err := sut.GetEntireExperience(uuid.New())

	// then
	assert.Equal(t, givenErr, err)
	assert.Nil(t, exp)
}

// ----- GetExperienceGainOfLessonCompletion -----
func TestGetExperienceGainReturnsNewLevelWhenUserLeveledUp(t *testing.T) {
	// given - a user which just completed 1 lesson
	m := repositories.NewMockLessonCompletionRepository(gomock.NewController(t))
	m.EXPECT().CountForUser(gomock.Any()).Return(1, nil)
	sut := ExperienceServiceImpl{LessonCompRepo: m}

	// when
	exp, err := sut.GetExperienceGainOfLessonCompletion(uuid.New())

	// then
	wantGain := domain_model.LessonCompletionGain
	assert.NoError(t, err)
	assert.Equal(t, wantGain, exp.NewExperienceGained)
	assert.Equal(t, wantGain, exp.TotalExperience)
	assert.Equal(t, 2, *exp.NewLevel)
}

func TestGetExperienceGainReturnsErrorFromRepository(t *testing.T) {
	// given
	m := repositories.NewMockLessonCompletionRepository(gomock.NewController(t))
	givenErr := errors.New("some error")
	m.EXPECT().CountForUser(gomock.Any()).Return(0, givenErr)
	sut := ExperienceServiceImpl{LessonCompRepo: m}

	// when
	exp, err := sut.GetExperienceGainOfLessonCompletion(uuid.New())

	// then
	assert.Equal(t, givenErr, err)
	assert.Nil(t, exp)
}

func calcLevel(totalExperience int) int {
	levelAsFloat := float64(1) + (math.Log((float64(totalExperience)/200)+1) / math.Log(1.5))
	lvl := int(math.Floor(levelAsFloat))
	return lvl
}
