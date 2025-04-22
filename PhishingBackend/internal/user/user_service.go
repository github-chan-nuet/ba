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
	Update(userId uuid.UUID, userPatchModel api.UserPatchModel) error
}

type ServiceImpl struct {
	Repo db.UserRepository
}

func (s *ServiceImpl) Update(userId uuid.UUID, userPatchModel api.UserPatchModel) error {
	user := &model.User{
		ID: userId,
	}
	if userPatchModel.Firstname != nil {
		user.Firstname = *userPatchModel.Firstname
	}
	if userPatchModel.Lastname != nil {
		user.Lastname = *userPatchModel.Lastname
	}
	if userPatchModel.Email != nil {
		user.Email = *userPatchModel.Email
	}
	if userPatchModel.Password != nil {
		hashedPw, err := HashPassword(*userPatchModel.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPw
	}
	return s.Repo.UpdateUser(user)
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
