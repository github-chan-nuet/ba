package web

import (
	"encoding/json"
	"net/http"
	"phishing_backend/api"
	"phishing_backend/internal/auth"
	"phishing_backend/internal/domain/db"
)

var authenticator auth.Authenticator = &auth.AuthenticatorImpl{
	UserRepository: &db.UserRepositoryImpl{},
}

func loginAndReturnJwtToken(w http.ResponseWriter, r *http.Request) {
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
