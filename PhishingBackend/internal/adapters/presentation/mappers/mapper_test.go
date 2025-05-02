package mappers

import (
	"github.com/stretchr/testify/assert"
	"phishing_backend/internal/adapters/presentation/api"
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

func ptr(s string) *string {
	return &s
}
