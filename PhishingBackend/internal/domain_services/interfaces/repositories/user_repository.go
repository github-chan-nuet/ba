package repositories

import (
	"errors"
	"phishing_backend/internal/domain_model"

	"github.com/google/uuid"
)

var ErrEmailAlreadyUsed = errors.New("email is already used")

type UserRepository interface {
	GetByEmailAndPassword(username string, password []byte) (*domain_model.User, error)
	CreateUser(user *domain_model.User) error
	GetUser(userId uuid.UUID) (*domain_model.User, error)
	UpdateUser(*domain_model.UserPatch) error
	GetAllUserIds() (*[]uuid.UUID, error)
}
