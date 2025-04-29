package repositories

import (
	"errors"
	"github.com/google/uuid"
	"phishing_backend/internal/domain_model"
)

var EmailAlreadyUsed = errors.New("email is already used")

type UserRepository interface {
	GetByEmailAndPassword(username string, password []byte) (*domain_model.User, error)
	CreateUser(user *domain_model.User) error
	GetUser(userId uuid.UUID) (*domain_model.User, error)
	UpdateUser(user *domain_model.User) error
}
