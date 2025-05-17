package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"phishing_backend/internal/domain_model"
	"testing"
)

func TestShouldCalculateCorrectScore(t *testing.T) {
	// given
	sut := ExamCompletionServiceImpl{}
	qs := []domain_model.ExamQuestion{
		{
			ID:       uuid.MustParse("00000000-0000-0000-0001-000000000000"),
			Question: "2 + 2 = ?",
			Answers: []domain_model.ExamQuestionAnswer{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0001-000000000001"),
					Answer:    "3",
					IsCorrect: false,
				},
				{
					ID:        uuid.MustParse("00000000-0000-0000-0001-000000000002"),
					Answer:    "4",
					IsCorrect: true,
				},
				{
					ID:        uuid.MustParse("00000000-0000-0000-0001-000000000003"),
					Answer:    "5 - 1",
					IsCorrect: true,
				},
			},
		},
	}
	exam := domain_model.Exam{Questions: qs}

	tests := []struct {
		name      string
		wantScore int
		answers   []uuid.UUID
	}{
		{"0 errors", 100, []uuid.UUID{
			uuid.MustParse("00000000-0000-0000-0001-000000000002"),
			uuid.MustParse("00000000-0000-0000-0001-000000000003"),
		}},
		{"1 error", 67, []uuid.UUID{
			uuid.MustParse("00000000-0000-0000-0001-000000000002"),
		}},
		{"2 errors", 33, []uuid.UUID{
			uuid.MustParse("00000000-0000-0000-0001-000000000001"),
			uuid.MustParse("00000000-0000-0000-0001-000000000002"),
		}},
		{"3 errors", 0, []uuid.UUID{
			uuid.MustParse("00000000-0000-0000-0001-000000000001"),
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			answers := []domain_model.QuestionCompletionDto{{
				Answers:    tt.answers,
				QuestionId: uuid.MustParse("00000000-0000-0000-0001-000000000000"),
			}}

			// when
			score, _ := sut.calculateScore(&exam, &answers)

			// then
			assert.Equal(t, tt.wantScore, score)
		})
	}
}

func TestShouldReturnErrorWhenQuestionDoesNotExist(t *testing.T) {
	// given
	sut := ExamCompletionServiceImpl{}
	qs := []domain_model.ExamQuestion{
		{
			ID:       uuid.MustParse("00000000-0000-0000-0001-000000000000"),
			Question: "2 + 2 = ?",
			Answers: []domain_model.ExamQuestionAnswer{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0001-000000000001"),
					Answer:    "3",
					IsCorrect: false,
				},
				{
					ID:        uuid.MustParse("00000000-0000-0000-0001-000000000002"),
					Answer:    "4",
					IsCorrect: true,
				},
			},
		},
	}
	exam := domain_model.Exam{Questions: qs}

	answers := []domain_model.QuestionCompletionDto{
		{
			Answers:    []uuid.UUID{uuid.MustParse("00000000-0000-0000-0001-000000000001")},
			QuestionId: uuid.MustParse("00000000-0000-0000-0001-000000000000"),
		},
		{
			Answers:    []uuid.UUID{uuid.MustParse("00000000-0000-0000-0001-000000000001")},
			QuestionId: uuid.MustParse("00000000-0000-0000-1111-000000000000"),
		},
	}

	// when
	_, err := sut.calculateScore(&exam, &answers)

	// then
	assert.EqualError(t, err, ErrQuestionNotExisting.Error())
}
