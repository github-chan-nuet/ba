package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"phishing_backend/internal/application/services"
	"phishing_backend/internal/infrastructure/persistance"
	"phishing_backend/internal/infrastructure/presentation/api"
)

var (
	userRepo                             = &persistance.UserRepositoryImpl{}
	authenticator services.Authenticator = &services.AuthenticatorImpl{UserRepository: userRepo}
	userService   services.UserService   = &services.UserServiceImpl{UserRepository: userRepo}
)

func LoginAndReturnJwtToken(w http.ResponseWriter, r *http.Request) {
	var auth api.UserAuthentication
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jwtToken, err := authenticator.Authenticate(auth.Email, auth.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jwtToken))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user api.UserPostModel
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = userService.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.PathValue("userId")
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authUser, err := authenticator.GetUser(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	if userId != authUser.ID {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	user, err := userService.Get(userId)
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
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
	authUser, err := authenticator.GetUser(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	if userId != authUser.ID {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	err = userService.Update(userId, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
