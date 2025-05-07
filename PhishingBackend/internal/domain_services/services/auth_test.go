package services

import (
	"errors"
	"fmt"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// ----- createJwtToken -----

func TestCreateValidJwtToken(t *testing.T) {
	// given
	sut := AuthenticatorImpl{}
	jwtKey = "3WVX3chChmoTVFjir5itbyKW5tnge7VY" // mock jwtKey
	user := domain_model.User{ID: uuid.MustParse("f5fe589e-481e-491b-914a-b0429ac842b4")}

	// when
	token, err := sut.createJwtToken(&user)

	// then
	assert.Nil(t, err)
	assert.NotNil(t, token)
	// and token is valid
	claims := validateTokenAndGetClaims(t, token)
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
		return []byte(jwtKey), nil
	})
	assert.NoError(t, err)
	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	return claims
}

// ----- GetUser -----

func TestReturnsErrorWhenHeaderIsEmpty(t *testing.T) {
	// given
	sut := AuthenticatorImpl{}

	// when
	token, err := sut.GetUser("")

	// then
	assert.Equal(t, uuid.Nil, token)
	assert.EqualError(t, err, "no JWT token present")
}

func TestReturnsErrorWhenHeaderIsInvalid(t *testing.T) {
	// given
	sut := AuthenticatorImpl{}
	tests := []struct {
		name   string
		header string
	}{
		{"no bearer", "Basic pw"},
		{"invalid structure", "fsfsfs"},
		{"no token", "Bearer"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// when
			token, err := sut.GetUser(tt.header)

			// then
			assert.Equal(t, uuid.Nil, token)
			assert.EqualError(t, err, "authorization header format must be: Bearer <token>")
		})
	}
}

func TestReturnsErrorWhenTokenIsInvalid(t *testing.T) {
	// given
	sut := AuthenticatorImpl{}

	// when
	token, err := sut.GetUser("Bearer <invalidstring>")

	// then
	assert.Equal(t, uuid.Nil, token)
	assert.Error(t, err)
}

func TestReturnsErrorWhenTokenSignatureIsInvalid(t *testing.T) {
	// given
	sut := AuthenticatorImpl{}
	jwtKey = "3WVX3chChmoTVFjir5itbyKW5tnge7VY" // mock jwtKey
	mockToken := createMockToken("random")

	// when
	token, err := sut.GetUser("Bearer " + mockToken)

	// then
	assert.Equal(t, uuid.Nil, token)
	assert.Error(t, err)
}

func TestReturnsErrorWhenTokenIsExpired(t *testing.T) {
	// given
	sut := AuthenticatorImpl{}
	key := "3WVX3chChmoTVFjir5itbyKW5tnge7VY"
	jwtKey = key // mock jwtKey
	mockToken := createMockToken(key)

	// when
	token, err := sut.GetUser("Bearer " + mockToken)

	// then
	assert.Equal(t, uuid.Nil, token)
	assert.EqualError(t, err, "invalid JWT token")
}

func createMockToken(key string) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  uuid.New(),
			"exp": jwt.NewNumericDate(time.Now().UTC().Add(-5 * time.Minute)),
		})
	signedToken, _ := t.SignedString([]byte(key))
	return signedToken
}

// ----- Authenticate -----

func TestReturnJwtTokenWhenUserExists(t *testing.T) {
	// given
	pbkdf2Iter = "1"
	ctrl := gomock.NewController(t)
	m := repositories.NewMockUserRepository(ctrl)
	m.EXPECT().GetByEmailAndPassword(gomock.Any(), gomock.Any()).Return(&domain_model.User{ID: uuid.New()}, nil)
	sut := AuthenticatorImpl{
		UserRepository: m,
	}

	// when
	jtwToken, err := sut.Authenticate("", "")

	// then
	assert.NoError(t, err)
	assert.NotNil(t, jtwToken)
}

func TestReturnErrorWhenUserDoesNotExist(t *testing.T) {
	// given
	pbkdf2Iter = "1"
	ctrl := gomock.NewController(t)
	m := repositories.NewMockUserRepository(ctrl)
	m.EXPECT().GetByEmailAndPassword(gomock.Any(), gomock.Any()).Return(nil, nil)
	sut := AuthenticatorImpl{
		UserRepository: m,
	}

	// when
	jtwToken, err := sut.Authenticate("", "")

	// then
	assert.Empty(t, jtwToken)
	assert.EqualError(t, err, "authentication claims are invalid")
}

func TestReturnErrorWhenRepositoryReturnsError(t *testing.T) {
	// given
	pbkdf2Iter = "1"
	ctrl := gomock.NewController(t)
	m := repositories.NewMockUserRepository(ctrl)
	m.EXPECT().GetByEmailAndPassword(gomock.Any(), gomock.Any()).Return(nil, errors.New("test error"))
	sut := AuthenticatorImpl{
		UserRepository: m,
	}

	// when
	jtwToken, err := sut.Authenticate("", "")

	// then
	assert.Empty(t, jtwToken)
	assert.EqualError(t, err, "test error")
}
