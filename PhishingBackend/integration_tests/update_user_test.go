//go:build integration

package integration_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"phishing_backend/internal/infrastructure/presentation/api"
	"testing"
)

func TestUserCanBeUpdated(t *testing.T) {
	// given
	user := createUser(t)
	jwtToken := getJwtTokenForUser(t, user)
	userId := getUserId(jwtToken)

	reqBody := api.UserAuthentication{
		Email:    user.Email,
		Password: user.Password,
	}
	marshal, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPatch, ts.URL+"/api/users/"+userId.String(), bytes.NewReader(marshal))

	// when
	resp, err := http.DefaultClient.Do(req)

	// then

}
