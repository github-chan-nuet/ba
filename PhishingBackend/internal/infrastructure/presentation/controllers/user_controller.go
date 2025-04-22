package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"phishing_backend/internal/application/services"
	"phishing_backend/internal/infrastructure/presentation/api"
)

type UserController struct {
	Authenticator services.Authenticator
	UserService   services.UserService
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
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
	err = c.UserService.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.PathValue("userId")
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authUser, err := c.Authenticator.GetUser(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	if userId != authUser.ID {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	user, err := c.UserService.Get(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	userDto := api.User{
		Email:           &user.Email,
		Firstname:       &user.Firstname,
		Lastname:        &user.Lastname,
		Level:           nil,
		TotalExperience: nil,
	}
	userJson, _ := json.Marshal(&userDto)
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
	authUser, err := c.Authenticator.GetUser(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	if userId != authUser.ID {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	err = c.UserService.Update(userId, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
