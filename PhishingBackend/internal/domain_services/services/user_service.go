package services

import (
	"github.com/google/uuid"
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
)

var _ UserService = (*UserServiceImpl)(nil)

type UserService interface {
	Create(userApiModel api.UserPostModel) error
	Get(userId uuid.UUID) (*domain_model.User, error)
	Update(userId uuid.UUID, userPatchModel api.UserPatchModel) error
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func (s *UserServiceImpl) Update(userId uuid.UUID, userPatchModel api.UserPatchModel) error {
	user := &domain_model.User{
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
	return s.UserRepository.UpdateUser(user)
}

func (s *UserServiceImpl) Get(userId uuid.UUID) (*domain_model.User, error) {
	return s.UserRepository.GetUser(userId)
}

func (s *UserServiceImpl) Create(userApiModel api.UserPostModel) error {
	hashedPw, err := HashPassword(userApiModel.Password)
	if err != nil {
		return err
	}
	user := &domain_model.User{
		ID:        uuid.New(),
		Firstname: userApiModel.Firstname,
		Lastname:  userApiModel.Lastname,
		Password:  hashedPw,
		Email:     userApiModel.Email,
	}
	err = s.UserRepository.CreateUser(user)
	return err
}
