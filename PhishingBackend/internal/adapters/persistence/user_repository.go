package persistence

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"log/slog"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
)

var _ repositories.UserRepository = (*UserRepositoryImpl)(nil)

const uniqueEmailConstraint = "users_email_key"

type UserRepositoryImpl struct {
}

func (u *UserRepositoryImpl) UpdateUser(user *domain_model.User) error {
	result := db.Model(user).Updates(*user)
	if result.Error != nil {
		slog.Error("Could not update user", "err", result.Error)
	}
	return result.Error
}

func (u *UserRepositoryImpl) GetUser(userId uuid.UUID) (*domain_model.User, error) {
	user := &domain_model.User{}
	result := db.First(&user, userId)
	if result.Error != nil {
		slog.Error("Could not get user by id", "err", result.Error)
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepositoryImpl) GetByEmailAndPassword(email string, password []byte) (*domain_model.User, error) {
	user := &domain_model.User{}
	result := db.Where("email = ? AND password = ?", email, password).First(user)
	if result.Error != nil {
		slog.Error("Could not get user by email and password", "err", result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepositoryImpl) CreateUser(user *domain_model.User) error {
	result := db.Create(user)
	if result.Error != nil {
		var e *pgconn.PgError
		if errors.As(result.Error, &e) {
			if e.Code == "23505" && e.ConstraintName == uniqueEmailConstraint {
				return repositories.ErrEmailAlreadyUsed
			}
		}
		slog.Error("Could not create user", "err", result.Error)
		return result.Error
	}
	return nil
}
