package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/adapters/presentation/error_handling"
	"phishing_backend/internal/adapters/presentation/mappers"
	"phishing_backend/internal/domain_services/services"
)

type UserController struct {
	Authenticator     services.Authenticator
	UserService       services.UserService
	ExperienceService services.ExperienceService
}

func (c *UserController) LoginAndReturnJwtToken(w http.ResponseWriter, r *http.Request) {
	var auth api.UserAuthentication
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jwtToken, err := c.Authenticator.Authenticate(auth.Email, auth.Password)
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jwtToken))
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user api.UserPostModel
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.UserService.Create(mappers.ToUserPostDto(user))
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.PathValue("userId")
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authUserId, err := c.Authenticator.GetUser(r.Header.Get("Authorization"))
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	if userId != authUserId {
		error_handling.WriteErrorDetailResponse(w, error_handling.ErrUnauthorized)
		return
	}
	user, err := c.UserService.Get(userId)
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	exp, err := c.ExperienceService.GetEntireExperience(userId)
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	userResp := api.User{
		Email:           &user.Email,
		Firstname:       &user.Firstname,
		Lastname:        &user.Lastname,
		Level:           &exp.Level,
		TotalExperience: &exp.TotalExperience,
	}
	userJson, _ := json.Marshal(&userResp)
	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.PathValue("userId")
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user api.UserPatchModel
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authUserId, err := c.Authenticator.GetUser(r.Header.Get("Authorization"))
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	if userId != authUserId {
		error_handling.WriteErrorDetailResponse(w, error_handling.ErrUnauthorized)
		return
	}
	err = c.UserService.Update(userId, mappers.ToUserPatchDto(user))
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
