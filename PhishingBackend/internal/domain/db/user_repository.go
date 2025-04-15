package db

import (
	"errors"
	"gorm.io/gorm"
	"log/slog"
	"phishing_backend/internal/domain/model"
)

var _ UserRepository = (*UserRepositoryImpl)(nil)

type UserRepository interface {
	GetByEmailAndPassword(username string, password []byte) (*model.User, error)
}

type UserRepositoryImpl struct {
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
