package domain_model

import (
	"github.com/google/uuid"
	"phishing_backend/internal/domain_model/validation"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Firstname string
	Lastname  string
	Password  []byte `gorm:"type:bytea"`
	Email     string
}

type UserPatchDto struct {
	Email     *string
	Firstname *string
	Lastname  *string
	Password  *string
}

func (u *UserPatchDto) Validate() error {
	if u.Email == nil && u.Firstname == nil && u.Lastname == nil && u.Password == nil {
		vErr := validation.NewValidationError()
		vErr.Add("#/", validation.NoFieldSet)
		return vErr
	}
	return nil
}

type UserPostDto struct {
	Email     string
	Firstname string
	Lastname  string
	Password  string
}

func (u *UserPostDto) Validate() error {
	vErr := validation.NewValidationError()
	if u.Firstname == "" {
		vErr.Add("#/firstname", validation.Mandatory)
	}
	if u.Lastname == "" {
		vErr.Add("#/lastname", validation.Mandatory)
	}
	if u.Email == "" {
		vErr.Add("#/email", validation.Mandatory)
	}
	if u.Password == "" {
		vErr.Add("#/password", validation.Mandatory)
	}
	if vErr.HasErr() {
		return vErr
	}
	return nil
}
