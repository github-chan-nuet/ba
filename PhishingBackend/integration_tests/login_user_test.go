//go:build integration

package integration_tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"phishing_backend/internal/adapters/presentation/api"
	"testing"
	"time"
)

func TestUserCanLogin(t *testing.T) {
	// given
	user := createUser(t)
	reqBody := api.UserAuthentication{
		Email:    user.Email,
		Password: user.Password,
	}
	marshal, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/users/login", bytes.NewReader(marshal))

	// when
	resp, err := http.DefaultClient.Do(req)

	// then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	// and token is valid
	binToken, _ := io.ReadAll(resp.Body)
	claims := validateTokenAndGetClaims(t, string(binToken))
	assert.NotNil(t, claims["id"])
	// and validity range is one week from approx. now
	iat, _ := claims.GetIssuedAt()
	assert.WithinDuration(t, time.Now(), iat.Time, 5*time.Minute)
	exp, _ := claims.GetExpirationTime()
	assert.WithinDuration(t, time.Now().Add(time.Hour*24*7), exp.Time, 5*time.Minute)
}

func validateTokenAndGetClaims(t *testing.T, rawToken string) jwt.MapClaims {
	token, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("PHBA_JWT_KEY")), nil
	})
	assert.NoError(t, err)
	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	return claims
}
