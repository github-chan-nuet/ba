package auth

import (
	"crypto/pbkdf2"
	"crypto/sha256"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"os"
	"phishing_backend/internal/domain/db"
	"phishing_backend/internal/domain/model"
	"strconv"
)

var (
	pbkdf2Salt = os.Getenv("PHBA_PBKDF2_SALT")
	pbkdf2Iter = os.Getenv("PHBA_PBKDF2_ITER")
	jwtKey     = os.Getenv("PHBA_JWT_KEY")
)

var _ Authenticator = (*AuthenticatorImpl)(nil)

type Authenticator interface {
	Authenticate(username, password string) (string, error)
}

type AuthenticatorImpl struct {
	UserRepository db.UserRepository
}

func (a *AuthenticatorImpl) Authenticate(username, password string) (string, error) {
	// https://de.wikipedia.org/wiki/PBKDF2
	hashedPassword, err := pbkdf2.Key(sha256.New, password, []byte(pbkdf2Salt), mustAtoi(pbkdf2Iter), 32)
	if err != nil {
		slog.Error("Could not hash password", "err", err)
		return "", err
	}
	user, err := a.UserRepository.GetByUsernameAndPassword(username, hashedPassword)
	if err != nil {
		return "", err
	}
	jwtToken, err := a.createJwtToken(user)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (a *AuthenticatorImpl) createJwtToken(user *model.User) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user.Name,
		})
	signedToken, err := t.SignedString([]byte(jwtKey))
	if err != nil {
		slog.Error("Could not sign JWT", "err", err)
		return "", err
	}
	return signedToken, nil
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("could not convert string to int, string: " + s)
	}
	return i
}
