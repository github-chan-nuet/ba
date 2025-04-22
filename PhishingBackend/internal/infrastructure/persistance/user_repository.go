package persistance

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log/slog"
	"phishing_backend/internal/application/interfaces/repositories"
	"phishing_backend/internal/domain"
)

var _ repositories.UserRepository = (*UserRepositoryImpl)(nil)

type UserRepositoryImpl struct {
}

func (u *UserRepositoryImpl) UpdateUser(user *domain.User) error {
	result := db.Save(user)
	if result.Error != nil {
		slog.Error("Could not update user", "err", result.Error)
	}
	return result.Error
}

func (u *UserRepositoryImpl) GetUser(userId uuid.UUID) (*domain.User, error) {
	user := &domain.User{}
	result := db.First(&user, userId)
	if result.Error != nil {
		slog.Error("Could not get user by id", "err", result.Error)
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepositoryImpl) GetByEmailAndPassword(email string, password []byte) (*domain.User, error) {
	user := &domain.User{}
	result := db.Where("email = ? AND password = ?", email, password).First(user)
	if result.Error != nil {
		slog.Error("Could not get user by email and password", "err", result.Error)
		if errors.As(result.Error, &gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepositoryImpl) CreateUser(user *domain.User) error {
	result := db.Create(user)
	if result.Error != nil {
		slog.Error("Could not create user", "err", result.Error)
		return result.Error
	}
	return nil
}
