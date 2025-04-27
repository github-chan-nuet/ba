//go:build integration

package integration_tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"phishing_backend/internal/application/services"
	"phishing_backend/internal/infrastructure/presentation/api"
	"testing"
)

func TestUserCanBeUpdated(t *testing.T) {
	// given
	user := createUser(t)
	jwtToken := getJwtTokenForUser(t, user)
	userId := getUserId(jwtToken)

	reqBody := api.UserPatchModel{
		Email:     ptr(createRandomEmail()),
		Password:  ptr("abc"),
		Firstname: ptr("d"),
		Lastname:  ptr("e"),
	}
	marshal, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPatch, ts.URL+"/api/users/"+userId.String(), bytes.NewReader(marshal))
	req.Header.Set("Authorization", "Bearer "+jwtToken)

	// when
	resp, err := http.DefaultClient.Do(req)

	// then
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	dbUser := getUser(*reqBody.Email)
	assert.Equal(t, *reqBody.Firstname, dbUser.Firstname)
	assert.Equal(t, *reqBody.Lastname, dbUser.Lastname)
	assert.Equal(t, *reqBody.Email, dbUser.Email)
	expectedPw, _ := services.HashPassword(*reqBody.Password)
	assert.Equal(t, expectedPw, dbUser.Password)
}

func ptr(s string) *string {
	return &s
}
