package persistence

import (
	"errors"
	"log/slog"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

var _ repositories.UserRepository = (*UserRepositoryImpl)(nil)

const uniqueEmailConstraint = "users_email_key"

type UserRepositoryImpl struct {
}

func (u *UserRepositoryImpl) UpdateUser(userPatch *domain_model.UserPatch) error {
	updates := make(map[string]interface{})

	if userPatch.Firstname != nil {
		updates["firstname"] = *userPatch.Firstname
	}
	if userPatch.Lastname != nil {
		updates["lastname"] = *userPatch.Lastname
	}
	if userPatch.Email != nil {
		updates["email"] = *userPatch.Email
	}
	if userPatch.Password != nil {
		updates["password"] = *userPatch.Password
	}
	if userPatch.ParticipatesInPhishingSimulation != nil {
		updates["participates_in_phishing_simulation"] = *userPatch.ParticipatesInPhishingSimulation
	}

	result := db.Model(&domain_model.User{ID: userPatch.ID}).Updates(updates)
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
	result := db.Where("email = ? AND password = ?", email, password).Limit(1).Find(user)
	if result.Error != nil {
		slog.Error("Could not get user by email and password", "err", result.Error)
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
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

func (u *UserRepositoryImpl) GetAllUsers() (*[]domain_model.User, error) {
	var users []domain_model.User
	result := db.Find(&users)
	if result.Error != nil {
		slog.Error("Could not get all users", "err", result.Error)
		return nil, result.Error
	}
	return &users, nil
}
