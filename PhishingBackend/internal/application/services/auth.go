package services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log/slog"
	"os"
	"phishing_backend/internal/application/interfaces/repositories"
	"phishing_backend/internal/domain"
	"strings"
	"time"
)

var (
	jwtKey               = os.Getenv("PHBA_JWT_KEY")
	_      Authenticator = (*AuthenticatorImpl)(nil)
)

const (
	jwtUserIdKey = "id"
	// https://www.rfc-editor.org/rfc/rfc7519#section-4.1.6
	jwtIssuedAtKey = "iat"
	// https://www.rfc-editor.org/rfc/rfc7519#section-4.1.4
	jwtExpirationTimeKey = "exp"
)

type Authenticator interface {
	Authenticate(username, password string) (string, error)
	GetUser(rawToken string) (uuid.UUID, error)
}

type AuthenticatorImpl struct {
	UserRepository repositories.UserRepository
}

func (a *AuthenticatorImpl) GetUser(authHeader string) (uuid.UUID, error) {
	if authHeader == "" {
		return uuid.Nil, errors.New("no token present")
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return uuid.Nil, fmt.Errorf("authorization header format must be: Bearer <token>")
	}
	claimsString := parts[1]
	// this code also verifies that the token is not expired
	token, err := jwt.Parse(claimsString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		fmt.Println("ahhhhhhhh")
		return uuid.Nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, errors.New("invalid token")
	}
	userId := claims[jwtUserIdKey]
	return uuid.MustParse(userId.(string)), nil
}

func (a *AuthenticatorImpl) Authenticate(email, password string) (string, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return "", err
	}
	user, err := a.UserRepository.GetByEmailAndPassword(email, hashedPassword)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("authentication claims are invalid")
	}
	jwtToken, err := a.createJwtToken(user)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (a *AuthenticatorImpl) createJwtToken(user *domain.User) (string, error) {
	currentTime := time.Now().UTC()
	inOneWeek := currentTime.AddDate(0, 0, 7)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			jwtUserIdKey:         user.ID.String(),
			jwtIssuedAtKey:       jwt.NewNumericDate(currentTime),
			jwtExpirationTimeKey: jwt.NewNumericDate(inOneWeek),
		})
	signedToken, err := t.SignedString([]byte(jwtKey))
	if err != nil {
		slog.Error("Could not sign JWT", "err", err)
		return "", err
	}
	return signedToken, nil
}
