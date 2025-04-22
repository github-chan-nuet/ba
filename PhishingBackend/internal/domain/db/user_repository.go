package db

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log/slog"
	"phishing_backend/internal/domain/model"
)

var _ UserRepository = (*UserRepositoryImpl)(nil)

type UserRepository interface {
	GetByEmailAndPassword(username string, password []byte) (*model.User, error)
	CreateUser(user *model.User) error
	GetUser(userId uuid.UUID) (*model.User, error)
	UpdateUser(user *model.User) error
}

type UserRepositoryImpl struct {
}

func (u *UserRepositoryImpl) UpdateUser(user *model.User) error {
	result := db.Save(user)
	if result.Error != nil {
		slog.Error("Could not update user", "err", result.Error)
	}
	return result.Error
}

func (u *UserRepositoryImpl) GetUser(userId uuid.UUID) (*model.User, error) {
	user := &model.User{}
	result := db.First(&user, userId)
	if result.Error != nil {
		slog.Error("Could not get user by id", "err", result.Error)
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepositoryImpl) GetByEmailAndPassword(email string, password []byte) (*model.User, error) {
	user := &model.User{}
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

func (u *UserRepositoryImpl) CreateUser(user *model.User) error {
	result := db.Create(user)
	if result.Error != nil {
		slog.Error("Could not create user", "err", result.Error)
		return result.Error
	}
	return nil
}
