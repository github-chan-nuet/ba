package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"testing"
	"time"
)

// ----- Create -----
func TestCreateCreatesNewLessonCompletion(t *testing.T) {
	// given
	m := repositories.NewMockLessonCompletionRepository(gomock.NewController(t))
	sut := LessonCompletionServiceImpl{Repo: m}
	var capture *domain_model.LessonCompletion
	m.EXPECT().Create(gomock.Any()).DoAndReturn(func(arg *domain_model.LessonCompletion) (int, error) {
		capture = arg
		return 1, nil
	})
	courseId, lessonId, userId := uuid.New(), uuid.New(), uuid.New()

	// when
	isNew, err := sut.Create(courseId, lessonId, userId)

	// then
	assert.NoError(t, err)
	assert.Equal(t, true, bool(isNew))
	assert.Equal(t, courseId, capture.CourseId)
	assert.Equal(t, lessonId, capture.LessonId)
	assert.Equal(t, userId, capture.UserFk)
	assert.NotEqual(t, uuid.Nil, capture.ID)
	assert.WithinDuration(t, time.Now().UTC(), capture.Time, 5*time.Second)
}

func TestCreateCreatesNoNewLessonWasCompleted(t *testing.T) {
	// given
	m := repositories.NewMockLessonCompletionRepository(gomock.NewController(t))
	sut := LessonCompletionServiceImpl{Repo: m}
	m.EXPECT().Create(gomock.Any()).Return(0, nil)
	courseId, lessonId, userId := uuid.New(), uuid.New(), uuid.New()

	// when
	isNew, err := sut.Create(courseId, lessonId, userId)

	// then
	assert.NoError(t, err)
	assert.Equal(t, false, bool(isNew))
}
