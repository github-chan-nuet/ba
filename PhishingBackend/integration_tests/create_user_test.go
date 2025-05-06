//go:build integration

package integration_tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/domain_services/services"
	"testing"
)

// https://medium.com/insiderengineering/integration-test-in-golang-899412b7e1bf
func TestNewUserCanBeCreated(t *testing.T) {
	// given
	reqBody := api.UserPostModel{
		Email:     "john.doe@test.com",
		Password:  "password",
		Firstname: "John",
		Lastname:  "Doe",
	}
	marshal, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/users", bytes.NewReader(marshal))

	// when
	resp, err := http.DefaultClient.Do(req)

	// then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	// and user was stored into the DB
	user := getUser(reqBody.Email)
	assert.NotNil(t, user)
	assert.Equal(t, reqBody.Email, user.Email)
	assert.Equal(t, reqBody.Firstname, user.Firstname)
	assert.Equal(t, reqBody.Lastname, user.Lastname)
	expectedPw := services.HashPassword(reqBody.Password)
	assert.Equal(t, expectedPw, user.Password)
}

func TestNewUserCantHaveSameEmail(t *testing.T) {
	// given
	reqBody := api.UserPostModel{
		Email:     createRandomEmail(),
		Password:  "password",
		Firstname: "John",
		Lastname:  "Doe",
	}
	marshal, _ := json.Marshal(reqBody)
	// a user exists
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/users", bytes.NewReader(marshal))
	_, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	req2, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/users", bytes.NewReader(marshal))

	// when
	resp, err := http.DefaultClient.Do(req2)

	// then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
	var gotProb api.ProblemDetail
	err = json.NewDecoder(resp.Body).Decode(&gotProb)
	assert.NoError(t, err)
	want := api.ProblemDetail{
		Type:   createErrorDetailUrn("email-already-used"),
		Title:  "Diese Email wird bereits verwendet",
		Status: 422,
	}
	assert.Equal(t, want, gotProb)
}
