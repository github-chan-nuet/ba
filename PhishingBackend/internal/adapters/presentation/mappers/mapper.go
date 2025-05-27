package mappers

import (
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/domain_model"
)

func ToUserPatchDto(patch api.UserPatchModel) *domain_model.UserPatchDto {
	return &domain_model.UserPatchDto{
		Email:                            patch.Email,
		Firstname:                        patch.Firstname,
		Lastname:                         patch.Lastname,
		Password:                         patch.Password,
		ParticipatesInPhishingSimulation: patch.ParticipatesInPhishingSimulation,
	}
}

func ToUserPostDto(post api.UserPostModel) *domain_model.UserPostDto {
	return &domain_model.UserPostDto{
		Email:                            post.Email,
		Firstname:                        post.Firstname,
		Lastname:                         post.Lastname,
		Password:                         post.Password,
		ParticipatesInPhishingSimulation: post.ParticipatesInPhishingSimulation,
	}
}

func ToQuestionCompletionDto(qs *[]api.QuestionCompletion) *[]domain_model.QuestionCompletionDto {
	dtos := make([]domain_model.QuestionCompletionDto, len(*qs))
	for i, q := range *qs {
		dtos[i] = domain_model.QuestionCompletionDto{
			Answers:    q.Answers,
			QuestionId: q.QuestionId,
		}
	}
	return &dtos
}

func ToApiExpGain(e *domain_model.ExperienceGain) *api.ExperienceGain {
	return &api.ExperienceGain{
		NewExperienceGained: e.NewExperienceGained,
		NewLevel:            e.NewLevel,
		TotalExperience:     e.TotalExperience,
	}
}
