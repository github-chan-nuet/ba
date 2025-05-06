package services

import (
	"github.com/google/uuid"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
)

var _ UserService = (*UserServiceImpl)(nil)

type UserService interface {
	Create(dto *domain_model.UserPostDto) error
	Get(userId uuid.UUID) (*domain_model.User, error)
	Update(userId uuid.UUID, dto *domain_model.UserPatchDto) error
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func (s *UserServiceImpl) Update(userId uuid.UUID, dto *domain_model.UserPatchDto) error {
	err := dto.Validate()
	if err != nil {
		return err
	}
	user := &domain_model.User{
		ID: userId,
	}
	if dto.Firstname != nil {
		user.Firstname = *dto.Firstname
	}
	if dto.Lastname != nil {
		user.Lastname = *dto.Lastname
	}
	if dto.Email != nil {
		user.Email = *dto.Email
	}
	if dto.Password != nil {
		hashedPw := HashPassword(*dto.Password)
		user.Password = hashedPw
	}
	return s.UserRepository.UpdateUser(user)
}

func (s *UserServiceImpl) Get(userId uuid.UUID) (*domain_model.User, error) {
	return s.UserRepository.GetUser(userId)
}

func (s *UserServiceImpl) Create(dto *domain_model.UserPostDto) error {
	err := dto.Validate()
	if err != nil {
		return err
	}
	hashedPw := HashPassword(dto.Password)
	user := &domain_model.User{
		ID:        uuid.New(),
		Firstname: dto.Firstname,
		Lastname:  dto.Lastname,
		Password:  hashedPw,
		Email:     dto.Email,
	}
	return s.UserRepository.CreateUser(user)
}
