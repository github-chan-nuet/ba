package mappers

import (
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/domain_model"
)

func ToUserPatchDto(patch api.UserPatchModel) *domain_model.UserPatchDto {
	return &domain_model.UserPatchDto{
		Email:     patch.Email,
		Firstname: patch.Firstname,
		Lastname:  patch.Lastname,
		Password:  patch.Password,
	}
}

func ToUserPostDto(post api.UserPostModel) *domain_model.UserPostDto {
	return &domain_model.UserPostDto{
		Email:     post.Email,
		Firstname: post.Firstname,
		Lastname:  post.Lastname,
		Password:  post.Password,
	}
}
