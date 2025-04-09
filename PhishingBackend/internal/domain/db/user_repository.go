package db

import (
	"log/slog"
	"phishing_backend/internal/domain/model"
)

var _ UserRepository = (*UserRepositoryImpl)(nil)

type UserRepository interface {
	GetByUsernameAndPassword(username string, password string) (*model.User, error)
}

type UserRepositoryImpl struct {
}

func (u *UserRepositoryImpl) GetByUsernameAndPassword(username string, password string) (*model.User, error) {
	user := &model.User{}
	result := db.Where("username = ? AND password = ?", username, password).First(user)
	if result.Error != nil {
		slog.Error("Could not get user by username and password", "err", result.Error)
		return nil, result.Error
	}
	return user, nil
}
