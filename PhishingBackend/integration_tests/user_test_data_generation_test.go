//go:build integration

package integration_tests

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"net/http"
	"phishing_backend/internal/infrastructure/presentation/api"
	"strings"
	"testing"
)

func createRandomEmail() string {
	return strings.ToLower(rand.Text()) + "@test.com"
}

func createUser(t *testing.T) *api.UserPostModel {
	reqBody := api.UserPostModel{
		Email:     createRandomEmail(),
		Password:  "password",
		Firstname: "John",
		Lastname:  "Doe",
	}
	marshal, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/users", bytes.NewReader(marshal))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("could not create user, expected %d, got %d", http.StatusCreated, resp.StatusCode)
		t.FailNow()
	}
	return &reqBody
}
