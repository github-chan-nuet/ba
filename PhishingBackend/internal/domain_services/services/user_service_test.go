package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"testing"
)

// ----- Update -----
func TestUpdateUpdatesUser(t *testing.T) {
	// given
	pbkdf2Iter = "1"
	m := repositories.NewMockUserRepository(gomock.NewController(t))
	var capture *domain_model.User
	m.EXPECT().UpdateUser(gomock.Any()).DoAndReturn(func(user *domain_model.User) error {
		capture = user
		return nil
	})

	sut := UserServiceImpl{UserRepository: m}
	userId := uuid.New()
	patch := domain_model.UserPatchDto{
		Email:     ptr("a@a"),
		Firstname: ptr("b"),
		Lastname:  ptr("c"),
		Password:  ptr("d"),
	}

	// when
	err := sut.Update(userId, &patch)

	// then
	assert.Nil(t, err)
	assert.Equal(t, *patch.Email, capture.Email)
	assert.Equal(t, *patch.Firstname, capture.Firstname)
	assert.Equal(t, *patch.Lastname, capture.Lastname)
	assert.Equal(t, userId, capture.ID)
	wantPw := HashPassword(*patch.Password)
	assert.Equal(t, wantPw, capture.Password)
}

func TestUpdateReturnsErrorOfRepository(t *testing.T) {
	// given
	m := repositories.NewMockUserRepository(gomock.NewController(t))
	wantErr := errors.New("repository error")
	m.EXPECT().UpdateUser(gomock.Any()).Return(wantErr)
	sut := UserServiceImpl{UserRepository: m}

	// when
	err := sut.Update(uuid.New(), &domain_model.UserPatchDto{})

	// then
	assert.Error(t, wantErr, err)
}

// ----- Create -----
func TestCreateCreatesUser(t *testing.T) {
	// given
	pbkdf2Iter = "1"
	m := repositories.NewMockUserRepository(gomock.NewController(t))
	var capture *domain_model.User
	m.EXPECT().CreateUser(gomock.Any()).DoAndReturn(func(user *domain_model.User) error {
		capture = user
		return nil
	})
	sut := UserServiceImpl{UserRepository: m}
	post := domain_model.UserPostDto{
		Email:     "a",
		Firstname: "b",
		Lastname:  "c",
		Password:  "d",
	}

	// when
	err := sut.Create(&post)

	// then
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, capture.ID)
	assert.Equal(t, post.Email, capture.Email)
	assert.Equal(t, post.Firstname, capture.Firstname)
	assert.Equal(t, post.Lastname, capture.Lastname)
	wantPw := HashPassword(post.Password)
	assert.Equal(t, wantPw, capture.Password)
}

func ptr(s string) *string {
	return &s
}
