package auth

import (
	"crypto/pbkdf2"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log/slog"
	"os"
	"phishing_backend/internal/domain/db"
	"phishing_backend/internal/domain/model"
	"strconv"
	"time"
)

var (
	pbkdf2Salt               = os.Getenv("PHBA_PBKDF2_SALT")
	pbkdf2Iter               = os.Getenv("PHBA_PBKDF2_ITER")
	jwtKey                   = os.Getenv("PHBA_JWT_KEY")
	_          Authenticator = (*AuthenticatorImpl)(nil)
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
	GetUser(rawToken string) (*model.User, error)
}

type AuthenticatorImpl struct {
	UserRepository db.UserRepository
}

func (a *AuthenticatorImpl) GetUser(rawToken string) (*model.User, error) {
	if rawToken == "" {
		return nil, errors.New("no token present")
	}
	token, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}
	exp, _ := claims.GetExpirationTime()
	if time.Now().UTC().After(exp.UTC()) {
		return nil, errors.New("token expired")
	}
	userId := claims[jwtUserIdKey]
	return &model.User{ID: uuid.MustParse(userId.(string))}, nil
}

func (a *AuthenticatorImpl) Authenticate(email, password string) (string, error) {
	// https://de.wikipedia.org/wiki/PBKDF2
	hashedPassword, err := pbkdf2.Key(sha256.New, password, []byte(pbkdf2Salt), mustAtoi(pbkdf2Iter), 32)
	if err != nil {
		slog.Error("Could not hash password", "err", err)
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

func (a *AuthenticatorImpl) createJwtToken(user *model.User) (string, error) {
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

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("could not convert string to int, string: " + s)
	}
	return i
}
