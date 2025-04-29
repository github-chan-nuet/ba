package repositories

import (
	"errors"
	"github.com/google/uuid"
	"phishing_backend/internal/domain"
)

var EmailAlreadyUsed = errors.New("email is already used")

type UserRepository interface {
	GetByEmailAndPassword(username string, password []byte) (*domain.User, error)
	CreateUser(user *domain.User) error
	GetUser(userId uuid.UUID) (*domain.User, error)
	UpdateUser(user *domain.User) error
}
