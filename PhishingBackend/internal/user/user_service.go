package user

import (
	"github.com/google/uuid"
	"phishing_backend/api"
	"phishing_backend/internal/domain/db"
	"phishing_backend/internal/domain/model"
)

var _ Service = (*ServiceImpl)(nil)

type Service interface {
	Create(userApiModel api.UserPostModel) error
	Get(userId uuid.UUID) (*model.User, error)
}

type ServiceImpl struct {
	Repo db.UserRepository
}

func (s *ServiceImpl) Get(userId uuid.UUID) (*model.User, error) {
	return s.Repo.GetUser(userId)
}

func (s *ServiceImpl) Create(userApiModel api.UserPostModel) error {
	hashedPw, err := HashPassword(userApiModel.Password)
	if err != nil {
		return err
	}
	user := &model.User{
		Firstname: userApiModel.Firstname,
		Lastname:  userApiModel.Lastname,
		Password:  hashedPw,
		Email:     userApiModel.Email,
	}
	err = s.Repo.CreateUser(user)
	return err
}
