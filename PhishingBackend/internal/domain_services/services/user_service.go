package services

import (
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"

	"github.com/google/uuid"
)

var _ UserService = (*UserServiceImpl)(nil)

type UserService interface {
	Create(dto *domain_model.UserPostDto) error
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

	userPatch := &domain_model.UserPatch{
		ID:                               userId,
		Firstname:                        dto.Firstname,
		Lastname:                         dto.Lastname,
		Email:                            dto.Email,
		ParticipatesInPhishingSimulation: dto.ParticipatesInPhishingSimulation,
	}

	if dto.Password != nil {
		hashedPw := HashPassword(*dto.Password)
		userPatch.Password = &hashedPw
	}

	return s.UserRepository.UpdateUser(userPatch)
}

func (s *UserServiceImpl) Create(dto *domain_model.UserPostDto) error {
	err := dto.Validate()
	if err != nil {
		return err
	}

	participatesInPhishingSimulation := false
	if dto.ParticipatesInPhishingSimulation != nil {
		participatesInPhishingSimulation = *dto.ParticipatesInPhishingSimulation
	}

	hashedPw := HashPassword(dto.Password)
	user := &domain_model.User{
		ID:                               uuid.New(),
		Firstname:                        dto.Firstname,
		Lastname:                         dto.Lastname,
		Password:                         hashedPw,
		Email:                            dto.Email,
		ParticipatesInPhishingSimulation: participatesInPhishingSimulation,
	}
	return s.UserRepository.CreateUser(user)
}
